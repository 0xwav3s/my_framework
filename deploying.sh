APP_NAME=$(basename $PWD)
export APP_NAME=$(echo $APP_NAME|sed -e 's/_/-/g'|cut -c1-7)
echo $APP_NAME

function build(){
    echo "Building Dockerfile..."
    docker build . -t protemon/$APP_NAME
    docker push protemon/$APP_NAME
    echo "Building Kube..."
    cd provision
    echo "Config mapp..."
    kubectl delete -f configmap.yaml
    kubectl apply -f configmap.yaml
    echo "Deploy on Kube..."
    sed 's/APP_NAME/'"$APP_NAME"'/g' k8s/* > zdeployment.yaml
    kubectl delete -f zdeployment.yaml
    kubectl apply -f zdeployment.yaml
}
build