///
//  Generated code. Do not modify.
//  source: global_entry_notify_api.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use errorDescriptor instead')
const Error$json = const {
  '1': 'Error',
  '2': const [
    const {'1': 'error', '3': 1, '4': 1, '5': 9, '10': 'error'},
  ],
};

/// Descriptor for `Error`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List errorDescriptor = $convert.base64Decode('CgVFcnJvchIUCgVlcnJvchgBIAEoCVIFZXJyb3I=');
@$core.Deprecated('Use updateNotificationTokenRqDescriptor instead')
const UpdateNotificationTokenRq$json = const {
  '1': 'UpdateNotificationTokenRq',
  '2': const [
    const {'1': 'userId', '3': 1, '4': 1, '5': 9, '10': 'userId'},
    const {'1': 'token', '3': 2, '4': 1, '5': 9, '10': 'token'},
  ],
};

/// Descriptor for `UpdateNotificationTokenRq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List updateNotificationTokenRqDescriptor = $convert.base64Decode('ChlVcGRhdGVOb3RpZmljYXRpb25Ub2tlblJxEhYKBnVzZXJJZBgBIAEoCVIGdXNlcklkEhQKBXRva2VuGAIgASgJUgV0b2tlbg==');
@$core.Deprecated('Use updateNotificationDetailsRqDescriptor instead')
const UpdateNotificationDetailsRq$json = const {
  '1': 'UpdateNotificationDetailsRq',
  '2': const [
    const {'1': 'userId', '3': 1, '4': 1, '5': 9, '10': 'userId'},
    const {'1': 'notificationDetails', '3': 2, '4': 1, '5': 11, '6': '.global_entry_notify_api.NotificationDetails', '10': 'notificationDetails'},
  ],
};

/// Descriptor for `UpdateNotificationDetailsRq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List updateNotificationDetailsRqDescriptor = $convert.base64Decode('ChtVcGRhdGVOb3RpZmljYXRpb25EZXRhaWxzUnESFgoGdXNlcklkGAEgASgJUgZ1c2VySWQSXgoTbm90aWZpY2F0aW9uRGV0YWlscxgCIAEoCzIsLmdsb2JhbF9lbnRyeV9ub3RpZnlfYXBpLk5vdGlmaWNhdGlvbkRldGFpbHNSE25vdGlmaWNhdGlvbkRldGFpbHM=');
@$core.Deprecated('Use deleteNotificationDetailsRqDescriptor instead')
const DeleteNotificationDetailsRq$json = const {
  '1': 'DeleteNotificationDetailsRq',
  '2': const [
    const {'1': 'userId', '3': 1, '4': 1, '5': 9, '10': 'userId'},
    const {'1': 'locationId', '3': 2, '4': 1, '5': 5, '10': 'locationId'},
  ],
};

/// Descriptor for `DeleteNotificationDetailsRq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List deleteNotificationDetailsRqDescriptor = $convert.base64Decode('ChtEZWxldGVOb3RpZmljYXRpb25EZXRhaWxzUnESFgoGdXNlcklkGAEgASgJUgZ1c2VySWQSHgoKbG9jYXRpb25JZBgCIAEoBVIKbG9jYXRpb25JZA==');
@$core.Deprecated('Use notificationDetailsDescriptor instead')
const NotificationDetails$json = const {
  '1': 'NotificationDetails',
  '2': const [
    const {'1': 'locationId', '3': 1, '4': 1, '5': 5, '10': 'locationId'},
    const {'1': 'targetDate', '3': 2, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'targetDate'},
    const {'1': 'appointmentDate', '3': 3, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'appointmentDate'},
    const {'1': 'lastNotifiedDate', '3': 4, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'lastNotifiedDate'},
  ],
};

/// Descriptor for `NotificationDetails`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List notificationDetailsDescriptor = $convert.base64Decode('ChNOb3RpZmljYXRpb25EZXRhaWxzEh4KCmxvY2F0aW9uSWQYASABKAVSCmxvY2F0aW9uSWQSOgoKdGFyZ2V0RGF0ZRgCIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXBSCnRhcmdldERhdGUSRAoPYXBwb2ludG1lbnREYXRlGAMgASgLMhouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcFIPYXBwb2ludG1lbnREYXRlEkYKEGxhc3ROb3RpZmllZERhdGUYBCABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wUhBsYXN0Tm90aWZpZWREYXRl');
@$core.Deprecated('Use locationsRpDescriptor instead')
const LocationsRp$json = const {
  '1': 'LocationsRp',
  '2': const [
    const {'1': 'locationId', '3': 1, '4': 1, '5': 5, '10': 'locationId'},
    const {'1': 'name', '3': 2, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'address', '3': 3, '4': 1, '5': 9, '10': 'address'},
    const {'1': 'addressAdditional', '3': 4, '4': 1, '5': 9, '10': 'addressAdditional'},
    const {'1': 'city', '3': 5, '4': 1, '5': 9, '10': 'city'},
    const {'1': 'state', '3': 6, '4': 1, '5': 9, '10': 'state'},
    const {'1': 'postalCode', '3': 7, '4': 1, '5': 9, '10': 'postalCode'},
    const {'1': 'countryCode', '3': 8, '4': 1, '5': 9, '10': 'countryCode'},
    const {'1': 'timezoneData', '3': 9, '4': 1, '5': 9, '10': 'timezoneData'},
    const {'1': 'phoneNumber', '3': 10, '4': 1, '5': 9, '10': 'phoneNumber'},
    const {'1': 'phoneAreaCode', '3': 11, '4': 1, '5': 9, '10': 'phoneAreaCode'},
    const {'1': 'phoneCountryCode', '3': 12, '4': 1, '5': 9, '10': 'phoneCountryCode'},
    const {'1': 'Directions', '3': 13, '4': 1, '5': 9, '10': 'Directions'},
    const {'1': 'Notes', '3': 14, '4': 1, '5': 9, '10': 'Notes'},
  ],
};

/// Descriptor for `LocationsRp`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List locationsRpDescriptor = $convert.base64Decode('CgtMb2NhdGlvbnNScBIeCgpsb2NhdGlvbklkGAEgASgFUgpsb2NhdGlvbklkEhIKBG5hbWUYAiABKAlSBG5hbWUSGAoHYWRkcmVzcxgDIAEoCVIHYWRkcmVzcxIsChFhZGRyZXNzQWRkaXRpb25hbBgEIAEoCVIRYWRkcmVzc0FkZGl0aW9uYWwSEgoEY2l0eRgFIAEoCVIEY2l0eRIUCgVzdGF0ZRgGIAEoCVIFc3RhdGUSHgoKcG9zdGFsQ29kZRgHIAEoCVIKcG9zdGFsQ29kZRIgCgtjb3VudHJ5Q29kZRgIIAEoCVILY291bnRyeUNvZGUSIgoMdGltZXpvbmVEYXRhGAkgASgJUgx0aW1lem9uZURhdGESIAoLcGhvbmVOdW1iZXIYCiABKAlSC3Bob25lTnVtYmVyEiQKDXBob25lQXJlYUNvZGUYCyABKAlSDXBob25lQXJlYUNvZGUSKgoQcGhvbmVDb3VudHJ5Q29kZRgMIAEoCVIQcGhvbmVDb3VudHJ5Q29kZRIeCgpEaXJlY3Rpb25zGA0gASgJUgpEaXJlY3Rpb25zEhQKBU5vdGVzGA4gASgJUgVOb3Rlcw==');
@$core.Deprecated('Use getNotificationDetailsRqDescriptor instead')
const GetNotificationDetailsRq$json = const {
  '1': 'GetNotificationDetailsRq',
  '2': const [
    const {'1': 'userId', '3': 1, '4': 1, '5': 9, '10': 'userId'},
  ],
};

/// Descriptor for `GetNotificationDetailsRq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getNotificationDetailsRqDescriptor = $convert.base64Decode('ChhHZXROb3RpZmljYXRpb25EZXRhaWxzUnESFgoGdXNlcklkGAEgASgJUgZ1c2VySWQ=');
@$core.Deprecated('Use getNotificationDetailsRpDescriptor instead')
const GetNotificationDetailsRp$json = const {
  '1': 'GetNotificationDetailsRp',
  '2': const [
    const {'1': 'notificationDetails', '3': 1, '4': 3, '5': 11, '6': '.global_entry_notify_api.NotificationDetails', '10': 'notificationDetails'},
  ],
};

/// Descriptor for `GetNotificationDetailsRp`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getNotificationDetailsRpDescriptor = $convert.base64Decode('ChhHZXROb3RpZmljYXRpb25EZXRhaWxzUnASXgoTbm90aWZpY2F0aW9uRGV0YWlscxgBIAMoCzIsLmdsb2JhbF9lbnRyeV9ub3RpZnlfYXBpLk5vdGlmaWNhdGlvbkRldGFpbHNSE25vdGlmaWNhdGlvbkRldGFpbHM=');
