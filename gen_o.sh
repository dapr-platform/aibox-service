dp-cli gen --connstr "postgresql://things:things2024@ali4:37104/thingsdb?sslmode=disable" \
--tables=o_aibox_device,o_aibox_event,o_aibox_update --model_naming "{{ toUpperCamelCase ( replace . \"o_\" \"\") }}"  \
--file_naming "{{ toLowerCamelCase ( replace . \"o_\" \"\") }}" \
--module aibox-service --api RUDB

