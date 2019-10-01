function build(){
    echo "Building Dockerfile..."
    docker build . -t protemon/my_framework
    docker push protemon/my_framework
    echo "Building Kube..."
    cd provision/k8s/
    kubectl apply -f deployment.yaml
    kubectl apply -f service.yaml
}
build