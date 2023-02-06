run:
	docker-compose up -d
	go build  &&  ./douyin_backend

init_api:
	hz new -mod douyin_backend -idl idl/core_api.thrift --model_dir=biz/hertz_gen/model
	make update_api

update_api:
	hz update -idl idl/core_api.thrift --model_dir=biz/hertz_gen/model
	hz update -idl idl/interact_api.thrift --model_dir=biz/hertz_gen/model
	hz update -idl idl/socialize_api.thrift --model_dir=biz/hertz_gen/model

