pipeline {
    agent any

    environment {
        DOCKER_PASSWORD = credentials('DOCKER_PASSWORD')
    }

    stages {
        stage('Init') {
            steps {
                echo 'Initializing..'
                echo "Running ${env.BUILD_ID} on ${env.JENKINS_URL}"
                echo "Current branch: ${env.BRANCH_NAME}"
                sh 'echo $DOCKER_PASSWORD | docker login -u sntshk --password-stdin'
            }
        }
        stage('Test') {
            steps {
                echo 'Testing..'
                // sh 'docker run --rm -e CI=true sntshk/cotu pytest'
            }
        }
        stage('Build, Publish & Cleanup: arm64') {
            steps {
                echo 'Building and publishing arm64 image to DockerHub..'
                sh 'GOOS=linux GOARCH=arm64 go build -buildvcs=false -o dist/cotu'
                sh 'docker buildx build --push --platform linux/arm64 -t sntshk/cotu:latest .'
                sh 'rm -rf dist/'
            }
        }
        stage('Build, Publish & Cleanup: amd64') {
            steps {
                echo 'Building and publishing amd64 image to DockerHub..'
                sh 'GOOS=linux GOARCH=amd64 go build -buildvcs=false -o dist/cotu'
                sh 'docker buildx build --push --platform linux/amd64 -t sntshk/cotu:latest .'
                sh 'rm -rf dist/'
            }
        }
        stage('Cleanup') {
            steps {
                echo 'Removing unused docker containers and images..'
                // keep intermediate images as cache, only delete the final image
                sh 'docker images -q sntshk/cotu | xargs --no-run-if-empty docker rmi'
            }
        }
    }
}
