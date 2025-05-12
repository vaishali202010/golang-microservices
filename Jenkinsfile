pipeline {
    agent any
    
    environment {
        DOCKER_HUB_CREDS = credentials('dockerhub-credentials')
        VERSION = "${env.BUILD_NUMBER}"
    }
    
    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        
        stage('Build and Test') {
            steps {
                sh '''
                for service in user-service product-service order-service payment-service inventory-service; do
                  cd $service
                  go mod tidy
                  go test ./... || true
                  cd ..
                done
                '''
            }
        }
        
        stage('Build and Push Docker Images') {
            steps {
                sh 'echo $DOCKER_HUB_CREDS_PSW | docker login -u $DOCKER_HUB_CREDS_USR --password-stdin'
                sh '''
                for service in user-service product-service order-service payment-service inventory-service; do
                  docker build -t $DOCKER_HUB_CREDS_USR/$service:$VERSION -t $DOCKER_HUB_CREDS_USR/$service:latest ./$service
                  docker push $DOCKER_HUB_CREDS_USR/$service:$VERSION
                  docker push $DOCKER_HUB_CREDS_USR/$service:latest
                done
                '''
            }
        }
        
        stage('Update Helm Chart') {
            steps {
                sh '''
                # Update image tags in values.yaml
                for service in user product order payment inventory; do
                  sed -i "s|repository: .*/$service-service|repository: $DOCKER_HUB_CREDS_USR/$service-service|g" helm/microservices/values.yaml
                  sed -i "s|tag: .*|tag: \\"$VERSION\\"|g" helm/microservices/values.yaml
                done
                
                # Commit and push changes to trigger ArgoCD sync
                git config user.email "jenkins@example.com"
                git config user.name "Jenkins CI"
                git add helm/microservices/values.yaml
                git commit -m "Update image versions to $VERSION [ci skip]" || echo "No changes to commit"
                git push origin HEAD:main
                '''
            }
        }
    }
    
    post {
        always {
            sh 'docker logout'
        }
    }
}