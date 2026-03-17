pipeline{
    agent any
    environment{
        PORT="8020"
        CONTAINER_NAME="go-app"
        IMAGE_NAME="ahmedkabil/go-app"
    }
    stages{
        stage("clean old container"){
            steps{
                sh "docker rm -f ${CONTAINER_NAME} || true"
            }
        }
        stage("build the app"){
            steps{
                sh "docker build --build-arg PORT=${PORT} -t ${IMAGE_NAME}:latest"
            }
        }
        stage("deploying the app"){
            steps{
                sh "docker run -d --name ${CONTAINER_NAME} -p ${PORT}:${PORT} ${IMAGE_NAME}:latest"
            }
        }
        stage("cleaning workspace"){
            steps{
                cleanWs()
            }
        }
    }
    post{
        success{
            echo "the app is running on port ${PORT}, http://192.168.50.134:${PORT}"
        }
    }
}