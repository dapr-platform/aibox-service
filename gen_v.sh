dp-cli gen --connstr "postgresql://things:things2024@ali4:37104/thingsdb?sslmode=disable" \
--tables=v_aibox_device_info,v_aibox_event_info,v_aibox_active_event_stats,v_aibox_update_info --model_naming "{{ toUpperCamelCase ( replace . \"v_\" \"\") }}"  \
--file_naming "{{ toLowerCamelCase ( replace . \"v_\" \"\") }}" \
--module aibox-service --api RUDB

