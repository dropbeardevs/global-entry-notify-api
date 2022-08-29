///
//  Generated code. Do not modify.
//  source: global_entry_notify_api.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'global_entry_notify_api.pb.dart' as $0;
import 'google/protobuf/empty.pb.dart' as $1;
export 'global_entry_notify_api.pb.dart';

class GlobalEntryNotifyServiceClient extends $grpc.Client {
  static final _$grpcUpdateNotificationToken = $grpc.ClientMethod<
          $0.UpdateNotificationTokenRq, $0.Error>(
      '/global_entry_notify_api.GlobalEntryNotifyService/GrpcUpdateNotificationToken',
      ($0.UpdateNotificationTokenRq value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Error.fromBuffer(value));
  static final _$grpcUpdateNotificationDetails = $grpc.ClientMethod<
          $0.UpdateNotificationDetailsRq, $0.Error>(
      '/global_entry_notify_api.GlobalEntryNotifyService/GrpcUpdateNotificationDetails',
      ($0.UpdateNotificationDetailsRq value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Error.fromBuffer(value));
  static final _$grpcDeleteNotificationDetails = $grpc.ClientMethod<
          $0.DeleteNotificationDetailsRq, $0.Error>(
      '/global_entry_notify_api.GlobalEntryNotifyService/GrpcDeleteNotificationDetails',
      ($0.DeleteNotificationDetailsRq value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Error.fromBuffer(value));
  static final _$grpcGetLocations =
      $grpc.ClientMethod<$1.Empty, $0.GetLocationsRp>(
          '/global_entry_notify_api.GlobalEntryNotifyService/GrpcGetLocations',
          ($1.Empty value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.GetLocationsRp.fromBuffer(value));
  static final _$grpcGetNotificationDetails = $grpc.ClientMethod<
          $0.GetNotificationDetailsRq, $0.GetNotificationDetailsRp>(
      '/global_entry_notify_api.GlobalEntryNotifyService/GrpcGetNotificationDetails',
      ($0.GetNotificationDetailsRq value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $0.GetNotificationDetailsRp.fromBuffer(value));

  GlobalEntryNotifyServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$0.Error> grpcUpdateNotificationToken(
      $0.UpdateNotificationTokenRq request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$grpcUpdateNotificationToken, request,
        options: options);
  }

  $grpc.ResponseFuture<$0.Error> grpcUpdateNotificationDetails(
      $0.UpdateNotificationDetailsRq request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$grpcUpdateNotificationDetails, request,
        options: options);
  }

  $grpc.ResponseFuture<$0.Error> grpcDeleteNotificationDetails(
      $0.DeleteNotificationDetailsRq request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$grpcDeleteNotificationDetails, request,
        options: options);
  }

  $grpc.ResponseFuture<$0.GetLocationsRp> grpcGetLocations($1.Empty request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$grpcGetLocations, request, options: options);
  }

  $grpc.ResponseFuture<$0.GetNotificationDetailsRp> grpcGetNotificationDetails(
      $0.GetNotificationDetailsRq request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$grpcGetNotificationDetails, request,
        options: options);
  }
}

abstract class GlobalEntryNotifyServiceBase extends $grpc.Service {
  $core.String get $name => 'global_entry_notify_api.GlobalEntryNotifyService';

  GlobalEntryNotifyServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.UpdateNotificationTokenRq, $0.Error>(
        'GrpcUpdateNotificationToken',
        grpcUpdateNotificationToken_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.UpdateNotificationTokenRq.fromBuffer(value),
        ($0.Error value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.UpdateNotificationDetailsRq, $0.Error>(
        'GrpcUpdateNotificationDetails',
        grpcUpdateNotificationDetails_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.UpdateNotificationDetailsRq.fromBuffer(value),
        ($0.Error value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.DeleteNotificationDetailsRq, $0.Error>(
        'GrpcDeleteNotificationDetails',
        grpcDeleteNotificationDetails_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.DeleteNotificationDetailsRq.fromBuffer(value),
        ($0.Error value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.Empty, $0.GetLocationsRp>(
        'GrpcGetLocations',
        grpcGetLocations_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.Empty.fromBuffer(value),
        ($0.GetLocationsRp value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.GetNotificationDetailsRq,
            $0.GetNotificationDetailsRp>(
        'GrpcGetNotificationDetails',
        grpcGetNotificationDetails_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.GetNotificationDetailsRq.fromBuffer(value),
        ($0.GetNotificationDetailsRp value) => value.writeToBuffer()));
  }

  $async.Future<$0.Error> grpcUpdateNotificationToken_Pre(
      $grpc.ServiceCall call,
      $async.Future<$0.UpdateNotificationTokenRq> request) async {
    return grpcUpdateNotificationToken(call, await request);
  }

  $async.Future<$0.Error> grpcUpdateNotificationDetails_Pre(
      $grpc.ServiceCall call,
      $async.Future<$0.UpdateNotificationDetailsRq> request) async {
    return grpcUpdateNotificationDetails(call, await request);
  }

  $async.Future<$0.Error> grpcDeleteNotificationDetails_Pre(
      $grpc.ServiceCall call,
      $async.Future<$0.DeleteNotificationDetailsRq> request) async {
    return grpcDeleteNotificationDetails(call, await request);
  }

  $async.Future<$0.GetLocationsRp> grpcGetLocations_Pre(
      $grpc.ServiceCall call, $async.Future<$1.Empty> request) async {
    return grpcGetLocations(call, await request);
  }

  $async.Future<$0.GetNotificationDetailsRp> grpcGetNotificationDetails_Pre(
      $grpc.ServiceCall call,
      $async.Future<$0.GetNotificationDetailsRq> request) async {
    return grpcGetNotificationDetails(call, await request);
  }

  $async.Future<$0.Error> grpcUpdateNotificationToken(
      $grpc.ServiceCall call, $0.UpdateNotificationTokenRq request);
  $async.Future<$0.Error> grpcUpdateNotificationDetails(
      $grpc.ServiceCall call, $0.UpdateNotificationDetailsRq request);
  $async.Future<$0.Error> grpcDeleteNotificationDetails(
      $grpc.ServiceCall call, $0.DeleteNotificationDetailsRq request);
  $async.Future<$0.GetLocationsRp> grpcGetLocations(
      $grpc.ServiceCall call, $1.Empty request);
  $async.Future<$0.GetNotificationDetailsRp> grpcGetNotificationDetails(
      $grpc.ServiceCall call, $0.GetNotificationDetailsRq request);
}
