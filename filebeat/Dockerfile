FROM docker.elastic.co/beats/filebeat-oss:7.15.0

COPY ./filebeat.yml /usr/share/filebeat/filebeat.yml

COPY ./logs/ /usr/share/filebeat/logs/

USER root
RUN chown root:filebeat /usr/share/filebeat/filebeat.yml
RUN chmod go-w /usr/share/filebeat/filebeat.yml

RUN chown -R root:filebeat /usr/share/filebeat/logs/
RUN chmod -R go-w /usr/share/filebeat/logs/
