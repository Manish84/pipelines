FROM continuumio/miniconda:4.7.12

RUN apt-get update --allow-releaseinfo-change && apt-get install -y --no-install-recommends \
    curl \
    wget \
    git \
    jq \
    tar \
    sudo

# Install eksctl
RUN curl --location "https://github.com/weaveworks/eksctl/releases/download/v0.113.0/eksctl_$(uname -s)_amd64.tar.gz" | tar xz -C /tmp \
 && mv /tmp/eksctl /usr/local/bin

# Install aws-iam-authenticator
RUN curl -S -o /usr/local/bin/aws-iam-authenticator https://amazon-eks.s3.us-west-2.amazonaws.com/1.16.8/2020-04-16/bin/linux/amd64/aws-iam-authenticator \
 && chmod +x /usr/local/bin/aws-iam-authenticator

# Install kubectl
RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.21.8/bin/linux/amd64/kubectl \
 && chmod +x ./kubectl \
 && cp ./kubectl /bin

# Install Argo CLI
RUN curl -sSL -o /usr/local/bin/argo https://github.com/argoproj/argo-workflows/releases/download/v2.8.0/argo-linux-amd64 \
 && chmod +x /usr/local/bin/argo

#Install helm(for ack)
RUN wget https://get.helm.sh/helm-v3.8.2-linux-amd64.tar.gz \ 
 && tar xvf helm-v3.8.2-linux-amd64.tar.gz && sudo mv linux-amd64/helm /usr/local/bin

#Install yq for ack
RUN wget https://github.com/mikefarah/yq/releases/download/v4.27.3/yq_linux_amd64 -O /usr/bin/yq &&\
    chmod +x /usr/bin/yq

# Copy conda environment early to avoid cache busting
COPY ./components/aws/sagemaker/tests/integration_tests/environment.yml environment.yml

# Create conda environment for running tests and set as start-up environment
RUN conda env create -f environment.yml
RUN echo "source activate kfp_test_env" > ~/.bashrc
ENV PATH "/opt/conda/envs/kfp_test_env/bin":$PATH

# Environment variables to be used by tests
ENV REGION="us-west-2"
ENV SAGEMAKER_EXECUTION_ROLE_ARN="arn:aws:iam::1234567890:role/sagemaker-role"
ENV ROBOMAKER_EXECUTION_ROLE_ARN="arn:aws:iam::1234567890:role/robomaker-role"
ENV S3_DATA_BUCKET="kfp-test-data"
ENV MINIO_LOCAL_PORT=9000
ENV KFP_NAMESPACE="kubeflow"

RUN mkdir pipelines
COPY ./ ./pipelines/

WORKDIR /pipelines/components/aws/sagemaker/tests/integration_tests/scripts/

ENTRYPOINT [ "/bin/bash", "./run_integration_tests" ]
