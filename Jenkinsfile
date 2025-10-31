pipeline {
    agent { label 'Built-In Node' }

    environment {
        APP_NAME = 'my-go-app.exe'
        APP_PORT = '8081'
    }

    stages {
        stage('Stop Existing Service') {
            steps {
                echo "Stopping any existing instance of ${env.APP_NAME}…"
                script {
                    // Only kill if the process is running, to avoid error "process not found"
                    bat """
                        tasklist /FI "IMAGENAME eq ${env.APP_NAME}" | find /I "${env.APP_NAME}" >nul
                        if %ERRORLEVEL%==0 (
                          echo Found ${env.APP_NAME}, killing process…
                          taskkill /F /IM ${env.APP_NAME}
                        ) else (
                          echo No running instance of ${env.APP_NAME} found.
                        )
                    """
                }
            }
        }

        stage('Build') {
            steps {
                echo 'Building Go application…'
                bat "go build -o ${env.APP_NAME}"
            }
        }

        stage('Unit Test') {
            steps {
                echo 'Running Go unit tests…'
                bat 'go test ./...'
            }
        }

        stage('Integration Test') {
            steps {
                echo 'Starting application for integration testing…'
                script {
                    // Launch app in background
                    bat "start /B ${env.APP_NAME} --port=${env.APP_PORT}"
                    echo 'Waiting 5 seconds for service to start…'
                    bat 'timeout /T 5 /NOBREAK'
                    echo 'Running integration tests…'
                    // Insert actual endpoint tests here, e.g. using powershell or curl
                    bat 'echo Integration tests succeeded (placeholder)'
                }
            }
        }
    }

    post {
        always {
            echo 'Post-build cleanup: ensuring the application is terminated.'
            script {
                bat """
                    tasklist /FI "IMAGENAME eq ${env.APP_NAME}" | find /I "${env.APP_NAME}" >nul
                    if %ERRORLEVEL%==0 (
                      echo ${env.APP_NAME} still running, killing process…
                      taskkill /F /IM ${env.APP_NAME}
                    ) else (
                      echo No running instance of ${env.APP_NAME} to clean up.
                    )
                """
            }
        }
    }
}
