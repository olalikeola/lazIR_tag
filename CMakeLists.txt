# Minimum CMake required
cmake_minimum_required(VERSION 3.15)

# Project
project(lazIR_tag)

# Protobuf
set(protobuf_MODULE_COMPATIBLE TRUE)
find_package(Protobuf CONFIG REQUIRED)
message(STATUS "Using protobuf ${Protobuf_VERSION}")

# Protobuf-compiler
set(_PROTOBUF_PROTOC $<TARGET_FILE:protobuf::protoc>)

# gRPC
find_package(gRPC CONFIG REQUIRED)
message(STATUS "Using gRPC ${gRPC_VERSION}")
set(_GRPC_GRPCPP gRPC::grpc++)
set(_GRPC_CPP_PLUGIN_EXECUTABLE $<TARGET_FILE:gRPC::grpc_cpp_plugin>)

# Proto file
get_filename_component(proto "src/proto/service.proto" ABSOLUTE)
get_filename_component(proto_path "${proto}" PATH)
message(STATUS "${proto}")
message(STATUS "${proto_path}")

# Generated sources
message(STATUS "${CMAKE_CURRENT_BINARY_DIR}")

set(proto_srcs "${proto_path}/../proto_c/service.pb.cc")
set(proto_hdrs "${proto_path}/../proto_c/service.pb.h")
set(grpc_srcs "${proto_path}/../proto_c/service.grpc.pb.cc")
set(grpc_hdrs "${proto_path}/../proto_c/service.grpc.pb.h")
add_custom_command(
      OUTPUT "${proto_srcs}" "${proto_hdrs}" "${grpc_srcs}" "${grpc_hdrs}"
      COMMAND ${_PROTOBUF_PROTOC}
      ARGS --grpc_out "${proto_path}/../proto_c"
        --cpp_out "${proto_path}/../proto_c"
        -I "${proto_path}"
        --plugin=protoc-gen-grpc="${_GRPC_CPP_PLUGIN_EXECUTABLE}"
        "${proto}"
      DEPENDS "${proto}")

# Include generated *.pb.h files
include_directories("${proto_path}")

# grpc_proto
add_library(grpc_proto
  ${grpc_srcs}
  ${grpc_hdrs}
  ${proto_srcs}
  ${proto_hdrs})
target_link_libraries(grpc_proto
  absl::check
  ${_REFLECTION}
  ${_GRPC_GRPCPP}
  ${_PROTOBUF_LIBPROTOBUF})
