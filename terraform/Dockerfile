FROM alpine:3.9

ARG ARG_AWS_ACCESS_KEY_ID
ARG ARG_AWS_SECRET_ACCESS_KEY
ARG ARG_DEFAULT_REGION

ENV AWS_ACCESS_KEY_ID=$ARG_AWS_ACCESS_KEY_ID
ENV AWS_SECRET_ACCESS_KEY=$ARG_AWS_SECRET_ACCESS_KEY
ENV AWS_DEFAULT_REGION=$ARG_DEFAULT_REGION

RUN apk add ca-certificates && rm -rf /var/cache/apk/*

COPY --from=hashicorp/terraform:light /bin/terraform /bin/

COPY . .

RUN terraform init

CMD ["terraform", "plan"]