# © 2018 Imhotep Software LLC. All materials licensed under Apache v2.0

# ----
# Build image
FROM   alpine:3.5
LABEL  maintainer="fernand@k8sland.io"

ADD execs /app

EXPOSE 4000
CMD    ["/app/depadm"]