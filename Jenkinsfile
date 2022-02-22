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
            }
        }
        stage('Test') {
            steps {
                echo 'Testing..'
                sh 'python3 -m pip install pytest'
                sh 'pytest'
            }
        }
        stage('Build') {
            steps {
                echo 'Running docker build -t sntshk/cotu:latest .'
                sh 'docker build -t sntshk/cotu:latest .'
            }
        }
        stage('Publish') {
            steps {
                echo 'Publishing..'
                echo 'Running docker push..'
            }
        }
        stage('Cleanup') {
            steps {
                echo 'docker rmi $(docker images -q)'
                sh 'docker rmi $(docker images -q)'
                sh 'rm -rf __pycache__'
                sh 'rm -rf .pytest_cache'
            }
        }
    }
}
