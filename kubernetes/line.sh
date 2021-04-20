#!/bin/bash

NAMESPACE=utility
FUNCTION=create

printHelp() {
  echo "Usage: "
  echo "  line.sh "
  echo 
  echo "  line.sh -h (print this message)"
  echo "  function : apply, create, delete"
  echo "Example :"
  echo "	line.sh -f create"
  echo
  exit 1
}

parseArguments() {
	while [ $# -gt 0 ]; do
		case $1 in
			--help | -h)
				printHelp
                ;;
            --function | -f)
                FUNCTION=$2
                ;;
		esac
		shift
	done
}

create() {
    echo "##########################################################"
    echo "############## Creating Line Notification ################"
    echo "##########################################################"

    kubectl create -f ./line-notification-configmap.yaml -n ${NAMESPACE}
    kubectl create -f ./line-notification-service.yaml -n ${NAMESPACE}
    kubectl create -f ./line-notification-deployment.yaml -n ${NAMESPACE}

    API_STATUS=$(kubectl get pods -n ${NAMESPACE} | grep line-notification-deployment | awk '{print $3}')
    while [ "${API_STATUS}" != "Running" ] && [ "${API_STATUS}" != "Running" ]; do
        if [ "${API_STATUS}" == "Error" ] || [ "${API_STATUS}" == "Error" ]; then
            echo "There is an error in api pod. Please check pod logs or describe."
            echo "- kubectl logs -f $(kubectl get pods -n ${NAMESPACE} | grep line-notification-deployment | awk '{print $1}') -n ${NAMESPACE}"
            echo "- kubectl describe pods $(kubectl get pods -n ${NAMESPACE} | grep line-notification-deployment | awk '{print $1}') -n ${NAMESPACE}"
            exit 1
        fi
        API_STATUS=$(kubectl get pods -n ${NAMESPACE} | grep line-notification-deployment | awk '{print $3}')
        echo "Waiting for pods to run. API Status = ${API_STATUS}"
        sleep 1
    done
}

delete() {
    echo "##########################################################"
    echo "############## Deleting Line Notification ################"
    echo "##########################################################"

    kubectl delete -f ./line-notification-deployment.yaml -n ${NAMESPACE}
    kubectl delete -f ./line-notification-service.yaml -n ${NAMESPACE}
    kubectl delete -f ./line-notification-configmap.yaml -n ${NAMESPACE}
}

apply() {
    echo "##########################################################"
    echo "############## Applying Line Notification ################"
    echo "##########################################################" 

    kubectl apply -f ./line-notification-configmap.yaml -n ${NAMESPACE}
    kubectl delete -f ./line-notification-deployment.yaml -n ${NAMESPACE}
    kubectl create -f ./line-notification-deployment.yaml -n ${NAMESPACE}

    API_STATUS=$(kubectl get pods -n ${NAMESPACE} | grep line-notification-deployment | awk '{print $3}')
    while [ "${API_STATUS}" != "Running" ] && [ "${API_STATUS}" != "Running" ]; do
        if [ "${API_STATUS}" == "Error" ] || [ "${API_STATUS}" == "Error" ]; then
            echo "There is an error in api pod. Please check pod logs or describe."
            echo "- kubectl logs -f $(kubectl get pods -n ${NAMESPACE} | grep line-notification-deployment | awk '{print $1}') -n ${NAMESPACE}"
            echo "- kubectl describe pods $(kubectl get pods -n ${NAMESPACE} | grep line-notification-deployment | awk '{print $1}') -n ${NAMESPACE}"
            exit 1
        fi
        API_STATUS=$(kubectl get pods -n ${NAMESPACE} | grep line-notification-deployment | awk '{print $3}')
        echo "Waiting for pods to run. API Status = ${API_STATUS}"
        sleep 1
    done
}

parseArguments $@

if [ ${FUNCTION} == "delete" ]; then
    delete
elif [ ${FUNCTION} == "create" ]; then
    create
elif [ ${FUNCTION} == "apply" ]; then
    apply
else
    printHelp
fi