pipeline {
    agent any

    environment {
        DOCKER_ID = credentials('DOCKER_ID')
        DOCKER_PASSWORD = credentials('DOCKER_PASSWORD')
    }

    stages {
        stage('Init') {
            steps {
                echo 'Initializing..'
                echo "Running ${env.BUILD_ID} on ${env.JENKINS_URL}"
                echo "Current branch: ${env.BRANCH_NAME}"
                sh 'echo $DOCKER_PASSWORD | docker login -u $DOCKER_ID --password-stdin'
            }
        }
        stage('Build') {
            steps {
                echo 'Running docker build -t sntshk/cotu:latest .'
                sh 'docker build -t sntshk/cotu:latest .'
            }
        }
        stage('Test') {
            steps {
                echo 'Testing..'
                sh 'docker run --rm -e CI=true sntshk/cotu pytest'
            }
        }
        stage('Publish') {
            steps {
                echo 'Publishing image to DockerHub..'
                sh 'docker push sntshk/cotu:latest'
            }
        }
        stage('Cleanup') {
            steps {
                echo 'docker rmi $(docker images -q)'
                sh 'docker rmi $(docker images -q)'
            }
        }
    }
}
