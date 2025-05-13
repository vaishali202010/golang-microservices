pipeline {
    agent any
    
    environment {
        DOCKER_HUB_CREDS = credentials('dockerhub-credentials')
    }
    
    stages {
        stage('Build and Test Go Code') {
            agent {
                docker {
                    image 'golang:1.22'
                    args '--network host -v /tmp/go-cache:/tmp/go-cache'
                    reuseNode true
                }
            }
            steps {
                sh '''
                # Set Go cache to a writable location
                export GOCACHE=/tmp/go-cache
                mkdir -p $GOCACHE
                
                ROOT_DIR=$(pwd)
                
                for service in user-service product-service order-service payment-service inventory-service; do
                  echo "Building and testing $service..."
                  cd "$ROOT_DIR/$service"
                  go mod tidy
                  go test ./... || true
                done
                '''
            }
        }
        
        stage('Build Docker Images') {
            steps {
                sh '''
                # Print diagnostic info
                echo "Current directory: $(pwd)"
                echo "Docker path: $(which docker)"
                
                for service in user-service product-service order-service payment-service inventory-service; do
                  echo "Building Docker image for $service..."
                  /var/lib/jenkins/jenkins-docker.sh build -t local/$service:latest ./$service
                done
                '''
            }
        }
        
        stage('Push Docker Images') {
            steps {
                sh '''
                echo $DOCKER_HUB_CREDS_PSW | /var/lib/jenkins/jenkins-docker.sh login -u $DOCKER_HUB_CREDS_USR --password-stdin
                
                for service in user-service product-service order-service payment-service inventory-service; do
                  echo "Tagging and pushing Docker image for $service..."
                  /var/lib/jenkins/jenkins-docker.sh tag local/$service:latest $DOCKER_HUB_CREDS_USR/$service:latest
                  /var/lib/jenkins/jenkins-docker.sh push $DOCKER_HUB_CREDS_USR/$service:latest
                done
                '''
            }
        }
    }
    
    post {
        always {
            sh '/var/lib/jenkins/jenkins-docker.sh logout || true'
            cleanWs()
        }
    }
}






















// pipeline {
//     agent {
//         docker {
//             image 'golang:1.22'
//             args '--network host'  // Use host network to ensure internet connectivity
//         }
//     }
    
//     environment {
//         DOCKER_HUB_CREDS = credentials('dockerhub-credentials')
//         GOCACHE = '/tmp/go-cache'
//     }
    
//     stages {
//         stage('Checkout') {
//             steps {
//                 checkout scm
//             }
//         }
        
//         stage('Build and Test') {
//             steps {
//                 sh '''
//                 # Store the root directory
//                 ROOT_DIR=$(pwd)
                
//                 for service in user-service product-service order-service payment-service inventory-service; do
//                   echo "Building and testing $service..."
//                   cd "$ROOT_DIR/$service"
//                   go mod tidy
//                   go test ./... || true
//                 done
//                 '''
//             }
//         }
        
//         stage('Build Docker Images') {
//             agent any  // Switch back to Jenkins agent for Docker operations
//             steps {
//                 sh '''
//                 for service in user-service product-service order-service payment-service inventory-service; do
//                   echo "Building Docker image for $service..."
//                   docker build -t local/$service:latest ./$service
//                 done
//                 '''
//             }
//         }
        
//         stage('Push Docker Images') {
//             agent any  // Use Jenkins agent for Docker operations
//             steps {
//                 sh 'echo $DOCKER_HUB_CREDS_PSW | docker login -u $DOCKER_HUB_CREDS_USR --password-stdin'
//                 sh '''
//                 for service in user-service product-service order-service payment-service inventory-service; do
//                   echo "Tagging and pushing Docker image for $service..."
//                   docker tag local/$service:latest $DOCKER_HUB_CREDS_USR/$service:latest
//                   docker push $DOCKER_HUB_CREDS_USR/$service:latest
//                 done
//                 '''
//             }
//         }
//     }
    
// post {
//     always {
//         script {
//             catchError(buildResult: 'SUCCESS', stageResult: 'SUCCESS') {
//                 sh 'docker logout || true'
//             }
//         }
//         cleanWs()
//     }
// }
// }




// pipeline {
//     agent any
    
//     environment {
//         DOCKER_HUB_CREDS = credentials('dockerhub-credentials')
//         VERSION = "${env.BUILD_NUMBER}"
//     }
    
//     stages {
//         stage('Checkout') {
//             steps {
//                 checkout scm
//             }
//         }
        
//         stage('Build and Test') {
//             steps {
//                 sh '''
//                 for service in user-service product-service order-service payment-service inventory-service; do
//                   cd $service
//                   go mod tidy
//                   go test ./... || true
//                   cd ..
//                 done
//                 '''
//             }
//         }
        
//         stage('Build and Push Docker Images') {
//             steps {
//                 sh 'echo $DOCKER_HUB_CREDS_PSW | docker login -u $DOCKER_HUB_CREDS_USR --password-stdin'
//                 sh '''
//                 for service in user-service product-service order-service payment-service inventory-service; do
//                   docker build -t $DOCKER_HUB_CREDS_USR/$service:$VERSION -t $DOCKER_HUB_CREDS_USR/$service:latest ./$service
//                   docker push $DOCKER_HUB_CREDS_USR/$service:$VERSION
//                   docker push $DOCKER_HUB_CREDS_USR/$service:latest
//                 done
//                 '''
//             }
//         }
        
//         stage('Update Helm Chart') {
//             steps {
//                 sh '''
//                 # Update image tags in values.yaml
//                 for service in user product order payment inventory; do
//                   sed -i "s|repository: .*/$service-service|repository: $DOCKER_HUB_CREDS_USR/$service-service|g" helm/microservices/values.yaml
//                   sed -i "s|tag: .*|tag: \\"$VERSION\\"|g" helm/microservices/values.yaml
//                 done
                
//                 # Commit and push changes to trigger ArgoCD sync
//                 git config user.email "jenkins@example.com"
//                 git config user.name "Jenkins CI"
//                 git add helm/microservices/values.yaml
//                 git commit -m "Update image versions to $VERSION [ci skip]" || echo "No changes to commit"
//                 git push origin HEAD:main
//                 '''
//             }
//         }
//     }
    
//     post {
//         always {
//             sh 'docker logout'
//         }
//     }
// }