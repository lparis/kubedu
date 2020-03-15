1. Install Docker Desktop for Mac from the Docker Store.

https://hub.docker.com/editions/community/docker-ce-desktop-mac

2. Enable Kubernetes.

- Click the Docker icon in the status bar. 
- Go to “Preferences”.
- Select the “Kubernetes” tab. 
- Check “Enable Kubernetes”.

3. Install or update kubectl CLI.

curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.15.5/bin/darwin/amd64/kubectl

4. Make the kubectl binary executable.

chmod +x ./kubectl

5. Move the binary in to your PATH.

sudo mv ./kubectl /usr/local/bin/kubectl

6. Test to ensure the version you installed is up-to-date:

kubectl version --client

