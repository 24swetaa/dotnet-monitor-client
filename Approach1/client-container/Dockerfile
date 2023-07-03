FROM python:3.9
EXPOSE 80
# Keeps Python from generating .pyc files in the container
ENV PYTHONDONTWRITEBYTECODE=1

# Turns off buffering for easier container logging
ENV PYTHONUNBUFFERED=1

# Install pip requirements
RUN pip install requests
RUN pip install kubernetes

WORKDIR /app
COPY . /app

# Creating environment variables
ARG ACTION
ENV ACTION=${ACTION:-"dump"}

ARG PODNAME
ENV PODNAME=${PODNAME:-"NO_PODNAME"}

ARG NAMESPACE
ENV NAMESPACE=${NAMESPACE:-"NAMESPACE"}

ARG UID
ENV UID=${UID:-"NO_UID"}

ARG PID
ENV PID=${PID:-"NO_PID"}

ARG NAME
ENV NAME=${NAME:-"NO_NAME"}

ARG DURATION
ENV DURATION=${DURATION:-"NO_DURATION"}

ARG EGRESS_PROVIDER
ENV EGRESS_PROVIDER=${EGRESS_PROVIDER:-"NO_EGRESS_PROVIDER"}

ARG TAG
ENV TAG=${TAG:-"NO_TAG"}

# Creates a non-root user with an explicit UID and adds permission to access the /app folder
# For more info, please refer to https://aka.ms/vscode-docker-python-configure-containers
RUN adduser -u 5678 --disabled-password --gecos "" appuser && chown -R appuser /app
USER appuser

# Modify permissions of the /app directory
RUN chmod -R 777 /app
# Switch to a non-root user
USER 1000

# During debugging, this entry point will be overridden. For more information, please refer to https://aka.ms/vscode-docker-python-debug
CMD ["sh", "-c" ,"python monitor_client.py --action $ACTION --pod-name $PODNAME --namespace $NAMESPACE --uid $UID --pid $PID --name $NAME --duration $DURATION --egressProvider $EGRESS_PROVIDER --tags $TAG "]
