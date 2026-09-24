// Jenkins declarative pipeline for building, testing, and releasing the provider.
pipeline {
    agent any

    environment {
        GO_VERSION = '1.23'
        GOFLAGS    = '-mod=mod'
        GOPROXY    = 'https://proxy.golang.org,direct'
    }

    options {
        timestamps()
        timeout(time: 60, unit: 'MINUTES')
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Setup Go') {
            steps {
                sh '''
                    set -euxo pipefail
                    if ! command -v go >/dev/null 2>&1; then
                        curl -sSfL "https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz" -o /tmp/go.tgz
                        sudo rm -rf /usr/local/go
                        sudo tar -C /usr/local -xzf /tmp/go.tgz
                    fi
                    export PATH="/usr/local/go/bin:${PATH}"
                    go version
                '''
            }
        }

        stage('Build') {
            steps {
                sh 'export PATH="/usr/local/go/bin:${PATH}"; go build -v ./...'
            }
        }

        stage('Vet') {
            steps {
                sh 'export PATH="/usr/local/go/bin:${PATH}"; go vet ./...'
            }
        }

        stage('Lint') {
            steps {
                sh '''
                    export PATH="/usr/local/go/bin:${PATH}"
                    ./pipelines/install_golangci-lint.sh
                    ./bin/golangci-lint run ./...
                '''
            }
        }

        stage('Test') {
            steps {
                sh '''
                    export PATH="/usr/local/go/bin:${PATH}"
                    go install github.com/jstemmer/go-junit-report/v2@latest
                    go test -v -coverprofile=coverage.txt ./... | tee test_output.txt
                    $(go env GOPATH)/bin/go-junit-report < test_output.txt > report.xml
                '''
            }
            post {
                always {
                    junit allowEmptyResults: true, testResults: 'report.xml'
                    archiveArtifacts artifacts: 'coverage.txt', allowEmptyArchive: true
                }
            }
        }

        // Optional: integration tests against a live tenant.
        // Secrets are bound only inside `withCredentials`, masked in logs, and
        // removed from the environment when the block exits — they are never
        // exposed to the whole stage. Store them in Jenkins Credentials Manager
        // as Secret Text: "example-host", "example-client-id", "example-client-secret".
        //
        // stage('Integration Tests') {
        //     when { branch 'main' }
        //     steps {
        //         withCredentials([
        //             string(credentialsId: 'example-host',          variable: 'EXAMPLE_HOST'),
        //             string(credentialsId: 'example-client-id',     variable: 'EXAMPLE_CLIENT_ID'),
        //             string(credentialsId: 'example-client-secret', variable: 'EXAMPLE_CLIENT_SECRET'),
        //         ]) {
        //             sh '''
        //                 export PATH="/usr/local/go/bin:${PATH}"
        //                 TF_ACC=1 go test -tags=integration -timeout 120m ./internal/...
        //             '''
        //         }
        //     }
        // }

        stage('Release') {
            when { tag 'v*' }
            steps {
                // Bind secrets only for this block. The GPG key is provided as a
                // *file* credential (its contents never touch an env var), and the
                // passphrase/token are masked strings. All are wiped when the block
                // ends. Store in Jenkins Credentials Manager:
                //   - gpg-private-key : Secret file (ASCII-armored private key)
                //   - gpg-passphrase  : Secret text
                //   - github-token    : Secret text
                withCredentials([
                    file(credentialsId: 'gpg-private-key', variable: 'GPG_KEY_FILE'),
                    string(credentialsId: 'gpg-passphrase', variable: 'GPG_PASSPHRASE'),
                    string(credentialsId: 'github-token', variable: 'GITHUB_TOKEN'),
                ]) {
                    sh '''
                        set -euo pipefail
                        export PATH="/usr/local/go/bin:${PATH}"
                        gpg --batch --import "$GPG_KEY_FILE"
                        export GPG_FINGERPRINT=$(gpg --list-secret-keys --with-colons | awk -F: '/^fpr:/ {print $10; exit}')
                        curl -sSfL https://goreleaser.com/static/run | bash -s -- release --clean
                    '''
                }
            }
        }
    }
}
