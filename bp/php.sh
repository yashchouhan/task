#!/bin/bash -ex

API_ENDPOINT=https://api.sys.test.cfdev.canopy-cloud.com
CF_ADMIN_USER=admin
CF_ADMIN_PASSWORD=mirorosyoydZorvucMop!FlevTyct1
CF_APP_DOMAIN=apps.test.cfdev.canopy-cloud.com
CF_SYS_DOMAIN=sys.test.cfdev.canopy-cloud.com

# Below Api Endpoint are need for CF Login Method
PREPROD_API_ENDPOINT=https://api.sys.preprod.cfdev.canopy-cloud.com
PROD_API_ENDPOINT=https://api.sys.eu01.cf.canopy-cloud.com



r=$RANDOM

#To Print Current Date and Time
now="$(date)"
printf "Current date and time %s\n" "$now"


Cf_Login()
{
 if [ "${API_ENDPOINT}" == "${PROD_API_ENDPOINT}" ]; then
   cf api $API_ENDPOINT
   cf auth $1 $2
 elif [ "${API_ENDPOINT}" == "${PREPROD_API_ENDPOINT}" ]; then
   cf api $API_ENDPOINT
   cf auth $1 $2
 else
   cf api --skip-ssl-validation $API_ENDPOINT
   cf auth $1 $2
 fi
}


Push_Php_Buildpack_App() {

  cd ../oscf-manifest-stubs/ci/smoketests/php

  APP_NAME="php-app-$r"
  echo "PHP Buildpack "
  cf buildpacks | egrep "php_buildpack"
  res=$?

  if [ $res -eq 0 ]; then
  echo "PHP buildpack found in cf buildpacks"
  else
  echo "PHP buildpack not found in CF buildpacks."
  exit 1
  fi

  cf push $APP_NAME -b php_buildpack
  APP_URL="http://$APP_NAME.$CF_APP_DOMAIN"
  curl $APP_URL
  cf apps
  cf services
  EXIT_STATUS=0
  cf env $APP_NAME
#  cf delete $APP_NAME -r -f
#  cf delete-orphaned-routes -f

}


set +e

Cf_Login $CF_ADMIN_USER $CF_ADMIN_PASSWORD
cf target -o oscf-sys -s testing
echo "Pushing app"
Push_Php_Buildpack_App
exit $EXIT_STATUS

