// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: service.proto
#ifndef GRPC_service_2eproto__INCLUDED
#define GRPC_service_2eproto__INCLUDED

#include "service.pb.h"

#include <functional>
#include <grpcpp/generic/async_generic_service.h>
#include <grpcpp/support/async_stream.h>
#include <grpcpp/support/async_unary_call.h>
#include <grpcpp/support/client_callback.h>
#include <grpcpp/client_context.h>
#include <grpcpp/completion_queue.h>
#include <grpcpp/support/message_allocator.h>
#include <grpcpp/support/method_handler.h>
#include <grpcpp/impl/proto_utils.h>
#include <grpcpp/impl/rpc_method.h>
#include <grpcpp/support/server_callback.h>
#include <grpcpp/impl/server_callback_handlers.h>
#include <grpcpp/server_context.h>
#include <grpcpp/impl/service_type.h>
#include <grpcpp/support/status.h>
#include <grpcpp/support/stub_options.h>
#include <grpcpp/support/sync_stream.h>

namespace lazIR_tag {

class Score final {
 public:
  static constexpr char const* service_full_name() {
    return "lazIR_tag.Score";
  }
  class StubInterface {
   public:
    virtual ~StubInterface() {}
    virtual ::grpc::Status AddRecord(::grpc::ClientContext* context, const ::lazIR_tag::AddRecordRequest& request, ::lazIR_tag::AddRecordResponse* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::lazIR_tag::AddRecordResponse>> AsyncAddRecord(::grpc::ClientContext* context, const ::lazIR_tag::AddRecordRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::lazIR_tag::AddRecordResponse>>(AsyncAddRecordRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::lazIR_tag::AddRecordResponse>> PrepareAsyncAddRecord(::grpc::ClientContext* context, const ::lazIR_tag::AddRecordRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::lazIR_tag::AddRecordResponse>>(PrepareAsyncAddRecordRaw(context, request, cq));
    }
    virtual ::grpc::Status GetRecord(::grpc::ClientContext* context, const ::lazIR_tag::GetRecordRequest& request, ::lazIR_tag::GetRecordResponse* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::lazIR_tag::GetRecordResponse>> AsyncGetRecord(::grpc::ClientContext* context, const ::lazIR_tag::GetRecordRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::lazIR_tag::GetRecordResponse>>(AsyncGetRecordRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::lazIR_tag::GetRecordResponse>> PrepareAsyncGetRecord(::grpc::ClientContext* context, const ::lazIR_tag::GetRecordRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::lazIR_tag::GetRecordResponse>>(PrepareAsyncGetRecordRaw(context, request, cq));
    }
    class async_interface {
     public:
      virtual ~async_interface() {}
      virtual void AddRecord(::grpc::ClientContext* context, const ::lazIR_tag::AddRecordRequest* request, ::lazIR_tag::AddRecordResponse* response, std::function<void(::grpc::Status)>) = 0;
      virtual void AddRecord(::grpc::ClientContext* context, const ::lazIR_tag::AddRecordRequest* request, ::lazIR_tag::AddRecordResponse* response, ::grpc::ClientUnaryReactor* reactor) = 0;
      virtual void GetRecord(::grpc::ClientContext* context, const ::lazIR_tag::GetRecordRequest* request, ::lazIR_tag::GetRecordResponse* response, std::function<void(::grpc::Status)>) = 0;
      virtual void GetRecord(::grpc::ClientContext* context, const ::lazIR_tag::GetRecordRequest* request, ::lazIR_tag::GetRecordResponse* response, ::grpc::ClientUnaryReactor* reactor) = 0;
    };
    typedef class async_interface experimental_async_interface;
    virtual class async_interface* async() { return nullptr; }
    class async_interface* experimental_async() { return async(); }
   private:
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::lazIR_tag::AddRecordResponse>* AsyncAddRecordRaw(::grpc::ClientContext* context, const ::lazIR_tag::AddRecordRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::lazIR_tag::AddRecordResponse>* PrepareAsyncAddRecordRaw(::grpc::ClientContext* context, const ::lazIR_tag::AddRecordRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::lazIR_tag::GetRecordResponse>* AsyncGetRecordRaw(::grpc::ClientContext* context, const ::lazIR_tag::GetRecordRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::lazIR_tag::GetRecordResponse>* PrepareAsyncGetRecordRaw(::grpc::ClientContext* context, const ::lazIR_tag::GetRecordRequest& request, ::grpc::CompletionQueue* cq) = 0;
  };
  class Stub final : public StubInterface {
   public:
    Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());
    ::grpc::Status AddRecord(::grpc::ClientContext* context, const ::lazIR_tag::AddRecordRequest& request, ::lazIR_tag::AddRecordResponse* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::lazIR_tag::AddRecordResponse>> AsyncAddRecord(::grpc::ClientContext* context, const ::lazIR_tag::AddRecordRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::lazIR_tag::AddRecordResponse>>(AsyncAddRecordRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::lazIR_tag::AddRecordResponse>> PrepareAsyncAddRecord(::grpc::ClientContext* context, const ::lazIR_tag::AddRecordRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::lazIR_tag::AddRecordResponse>>(PrepareAsyncAddRecordRaw(context, request, cq));
    }
    ::grpc::Status GetRecord(::grpc::ClientContext* context, const ::lazIR_tag::GetRecordRequest& request, ::lazIR_tag::GetRecordResponse* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::lazIR_tag::GetRecordResponse>> AsyncGetRecord(::grpc::ClientContext* context, const ::lazIR_tag::GetRecordRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::lazIR_tag::GetRecordResponse>>(AsyncGetRecordRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::lazIR_tag::GetRecordResponse>> PrepareAsyncGetRecord(::grpc::ClientContext* context, const ::lazIR_tag::GetRecordRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::lazIR_tag::GetRecordResponse>>(PrepareAsyncGetRecordRaw(context, request, cq));
    }
    class async final :
      public StubInterface::async_interface {
     public:
      void AddRecord(::grpc::ClientContext* context, const ::lazIR_tag::AddRecordRequest* request, ::lazIR_tag::AddRecordResponse* response, std::function<void(::grpc::Status)>) override;
      void AddRecord(::grpc::ClientContext* context, const ::lazIR_tag::AddRecordRequest* request, ::lazIR_tag::AddRecordResponse* response, ::grpc::ClientUnaryReactor* reactor) override;
      void GetRecord(::grpc::ClientContext* context, const ::lazIR_tag::GetRecordRequest* request, ::lazIR_tag::GetRecordResponse* response, std::function<void(::grpc::Status)>) override;
      void GetRecord(::grpc::ClientContext* context, const ::lazIR_tag::GetRecordRequest* request, ::lazIR_tag::GetRecordResponse* response, ::grpc::ClientUnaryReactor* reactor) override;
     private:
      friend class Stub;
      explicit async(Stub* stub): stub_(stub) { }
      Stub* stub() { return stub_; }
      Stub* stub_;
    };
    class async* async() override { return &async_stub_; }

   private:
    std::shared_ptr< ::grpc::ChannelInterface> channel_;
    class async async_stub_{this};
    ::grpc::ClientAsyncResponseReader< ::lazIR_tag::AddRecordResponse>* AsyncAddRecordRaw(::grpc::ClientContext* context, const ::lazIR_tag::AddRecordRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::lazIR_tag::AddRecordResponse>* PrepareAsyncAddRecordRaw(::grpc::ClientContext* context, const ::lazIR_tag::AddRecordRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::lazIR_tag::GetRecordResponse>* AsyncGetRecordRaw(::grpc::ClientContext* context, const ::lazIR_tag::GetRecordRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::lazIR_tag::GetRecordResponse>* PrepareAsyncGetRecordRaw(::grpc::ClientContext* context, const ::lazIR_tag::GetRecordRequest& request, ::grpc::CompletionQueue* cq) override;
    const ::grpc::internal::RpcMethod rpcmethod_AddRecord_;
    const ::grpc::internal::RpcMethod rpcmethod_GetRecord_;
  };
  static std::unique_ptr<Stub> NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());

  class Service : public ::grpc::Service {
   public:
    Service();
    virtual ~Service();
    virtual ::grpc::Status AddRecord(::grpc::ServerContext* context, const ::lazIR_tag::AddRecordRequest* request, ::lazIR_tag::AddRecordResponse* response);
    virtual ::grpc::Status GetRecord(::grpc::ServerContext* context, const ::lazIR_tag::GetRecordRequest* request, ::lazIR_tag::GetRecordResponse* response);
  };
  template <class BaseClass>
  class WithAsyncMethod_AddRecord : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithAsyncMethod_AddRecord() {
      ::grpc::Service::MarkMethodAsync(0);
    }
    ~WithAsyncMethod_AddRecord() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status AddRecord(::grpc::ServerContext* /*context*/, const ::lazIR_tag::AddRecordRequest* /*request*/, ::lazIR_tag::AddRecordResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestAddRecord(::grpc::ServerContext* context, ::lazIR_tag::AddRecordRequest* request, ::grpc::ServerAsyncResponseWriter< ::lazIR_tag::AddRecordResponse>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithAsyncMethod_GetRecord : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithAsyncMethod_GetRecord() {
      ::grpc::Service::MarkMethodAsync(1);
    }
    ~WithAsyncMethod_GetRecord() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetRecord(::grpc::ServerContext* /*context*/, const ::lazIR_tag::GetRecordRequest* /*request*/, ::lazIR_tag::GetRecordResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestGetRecord(::grpc::ServerContext* context, ::lazIR_tag::GetRecordRequest* request, ::grpc::ServerAsyncResponseWriter< ::lazIR_tag::GetRecordResponse>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(1, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  typedef WithAsyncMethod_AddRecord<WithAsyncMethod_GetRecord<Service > > AsyncService;
  template <class BaseClass>
  class WithCallbackMethod_AddRecord : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithCallbackMethod_AddRecord() {
      ::grpc::Service::MarkMethodCallback(0,
          new ::grpc::internal::CallbackUnaryHandler< ::lazIR_tag::AddRecordRequest, ::lazIR_tag::AddRecordResponse>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::lazIR_tag::AddRecordRequest* request, ::lazIR_tag::AddRecordResponse* response) { return this->AddRecord(context, request, response); }));}
    void SetMessageAllocatorFor_AddRecord(
        ::grpc::MessageAllocator< ::lazIR_tag::AddRecordRequest, ::lazIR_tag::AddRecordResponse>* allocator) {
      ::grpc::internal::MethodHandler* const handler = ::grpc::Service::GetHandler(0);
      static_cast<::grpc::internal::CallbackUnaryHandler< ::lazIR_tag::AddRecordRequest, ::lazIR_tag::AddRecordResponse>*>(handler)
              ->SetMessageAllocator(allocator);
    }
    ~WithCallbackMethod_AddRecord() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status AddRecord(::grpc::ServerContext* /*context*/, const ::lazIR_tag::AddRecordRequest* /*request*/, ::lazIR_tag::AddRecordResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* AddRecord(
      ::grpc::CallbackServerContext* /*context*/, const ::lazIR_tag::AddRecordRequest* /*request*/, ::lazIR_tag::AddRecordResponse* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithCallbackMethod_GetRecord : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithCallbackMethod_GetRecord() {
      ::grpc::Service::MarkMethodCallback(1,
          new ::grpc::internal::CallbackUnaryHandler< ::lazIR_tag::GetRecordRequest, ::lazIR_tag::GetRecordResponse>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::lazIR_tag::GetRecordRequest* request, ::lazIR_tag::GetRecordResponse* response) { return this->GetRecord(context, request, response); }));}
    void SetMessageAllocatorFor_GetRecord(
        ::grpc::MessageAllocator< ::lazIR_tag::GetRecordRequest, ::lazIR_tag::GetRecordResponse>* allocator) {
      ::grpc::internal::MethodHandler* const handler = ::grpc::Service::GetHandler(1);
      static_cast<::grpc::internal::CallbackUnaryHandler< ::lazIR_tag::GetRecordRequest, ::lazIR_tag::GetRecordResponse>*>(handler)
              ->SetMessageAllocator(allocator);
    }
    ~WithCallbackMethod_GetRecord() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetRecord(::grpc::ServerContext* /*context*/, const ::lazIR_tag::GetRecordRequest* /*request*/, ::lazIR_tag::GetRecordResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* GetRecord(
      ::grpc::CallbackServerContext* /*context*/, const ::lazIR_tag::GetRecordRequest* /*request*/, ::lazIR_tag::GetRecordResponse* /*response*/)  { return nullptr; }
  };
  typedef WithCallbackMethod_AddRecord<WithCallbackMethod_GetRecord<Service > > CallbackService;
  typedef CallbackService ExperimentalCallbackService;
  template <class BaseClass>
  class WithGenericMethod_AddRecord : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithGenericMethod_AddRecord() {
      ::grpc::Service::MarkMethodGeneric(0);
    }
    ~WithGenericMethod_AddRecord() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status AddRecord(::grpc::ServerContext* /*context*/, const ::lazIR_tag::AddRecordRequest* /*request*/, ::lazIR_tag::AddRecordResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithGenericMethod_GetRecord : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithGenericMethod_GetRecord() {
      ::grpc::Service::MarkMethodGeneric(1);
    }
    ~WithGenericMethod_GetRecord() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetRecord(::grpc::ServerContext* /*context*/, const ::lazIR_tag::GetRecordRequest* /*request*/, ::lazIR_tag::GetRecordResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithRawMethod_AddRecord : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawMethod_AddRecord() {
      ::grpc::Service::MarkMethodRaw(0);
    }
    ~WithRawMethod_AddRecord() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status AddRecord(::grpc::ServerContext* /*context*/, const ::lazIR_tag::AddRecordRequest* /*request*/, ::lazIR_tag::AddRecordResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestAddRecord(::grpc::ServerContext* context, ::grpc::ByteBuffer* request, ::grpc::ServerAsyncResponseWriter< ::grpc::ByteBuffer>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithRawMethod_GetRecord : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawMethod_GetRecord() {
      ::grpc::Service::MarkMethodRaw(1);
    }
    ~WithRawMethod_GetRecord() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetRecord(::grpc::ServerContext* /*context*/, const ::lazIR_tag::GetRecordRequest* /*request*/, ::lazIR_tag::GetRecordResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestGetRecord(::grpc::ServerContext* context, ::grpc::ByteBuffer* request, ::grpc::ServerAsyncResponseWriter< ::grpc::ByteBuffer>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(1, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithRawCallbackMethod_AddRecord : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawCallbackMethod_AddRecord() {
      ::grpc::Service::MarkMethodRawCallback(0,
          new ::grpc::internal::CallbackUnaryHandler< ::grpc::ByteBuffer, ::grpc::ByteBuffer>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::grpc::ByteBuffer* request, ::grpc::ByteBuffer* response) { return this->AddRecord(context, request, response); }));
    }
    ~WithRawCallbackMethod_AddRecord() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status AddRecord(::grpc::ServerContext* /*context*/, const ::lazIR_tag::AddRecordRequest* /*request*/, ::lazIR_tag::AddRecordResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* AddRecord(
      ::grpc::CallbackServerContext* /*context*/, const ::grpc::ByteBuffer* /*request*/, ::grpc::ByteBuffer* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithRawCallbackMethod_GetRecord : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawCallbackMethod_GetRecord() {
      ::grpc::Service::MarkMethodRawCallback(1,
          new ::grpc::internal::CallbackUnaryHandler< ::grpc::ByteBuffer, ::grpc::ByteBuffer>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::grpc::ByteBuffer* request, ::grpc::ByteBuffer* response) { return this->GetRecord(context, request, response); }));
    }
    ~WithRawCallbackMethod_GetRecord() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetRecord(::grpc::ServerContext* /*context*/, const ::lazIR_tag::GetRecordRequest* /*request*/, ::lazIR_tag::GetRecordResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* GetRecord(
      ::grpc::CallbackServerContext* /*context*/, const ::grpc::ByteBuffer* /*request*/, ::grpc::ByteBuffer* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_AddRecord : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithStreamedUnaryMethod_AddRecord() {
      ::grpc::Service::MarkMethodStreamed(0,
        new ::grpc::internal::StreamedUnaryHandler<
          ::lazIR_tag::AddRecordRequest, ::lazIR_tag::AddRecordResponse>(
            [this](::grpc::ServerContext* context,
                   ::grpc::ServerUnaryStreamer<
                     ::lazIR_tag::AddRecordRequest, ::lazIR_tag::AddRecordResponse>* streamer) {
                       return this->StreamedAddRecord(context,
                         streamer);
                  }));
    }
    ~WithStreamedUnaryMethod_AddRecord() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status AddRecord(::grpc::ServerContext* /*context*/, const ::lazIR_tag::AddRecordRequest* /*request*/, ::lazIR_tag::AddRecordResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedAddRecord(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::lazIR_tag::AddRecordRequest,::lazIR_tag::AddRecordResponse>* server_unary_streamer) = 0;
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_GetRecord : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithStreamedUnaryMethod_GetRecord() {
      ::grpc::Service::MarkMethodStreamed(1,
        new ::grpc::internal::StreamedUnaryHandler<
          ::lazIR_tag::GetRecordRequest, ::lazIR_tag::GetRecordResponse>(
            [this](::grpc::ServerContext* context,
                   ::grpc::ServerUnaryStreamer<
                     ::lazIR_tag::GetRecordRequest, ::lazIR_tag::GetRecordResponse>* streamer) {
                       return this->StreamedGetRecord(context,
                         streamer);
                  }));
    }
    ~WithStreamedUnaryMethod_GetRecord() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status GetRecord(::grpc::ServerContext* /*context*/, const ::lazIR_tag::GetRecordRequest* /*request*/, ::lazIR_tag::GetRecordResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedGetRecord(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::lazIR_tag::GetRecordRequest,::lazIR_tag::GetRecordResponse>* server_unary_streamer) = 0;
  };
  typedef WithStreamedUnaryMethod_AddRecord<WithStreamedUnaryMethod_GetRecord<Service > > StreamedUnaryService;
  typedef Service SplitStreamedService;
  typedef WithStreamedUnaryMethod_AddRecord<WithStreamedUnaryMethod_GetRecord<Service > > StreamedService;
};

}  // namespace lazIR_tag


#endif  // GRPC_service_2eproto__INCLUDED