///
//  Generated code. Do not modify.
//  source: global_entry_notify_api.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'google/protobuf/timestamp.pb.dart' as $2;

class Error extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Error', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'global_entry_notify_api'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'error')
    ..hasRequiredFields = false
  ;

  Error._() : super();
  factory Error({
    $core.String? error,
  }) {
    final _result = create();
    if (error != null) {
      _result.error = error;
    }
    return _result;
  }
  factory Error.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Error.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Error clone() => Error()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Error copyWith(void Function(Error) updates) => super.copyWith((message) => updates(message as Error)) as Error; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Error create() => Error._();
  Error createEmptyInstance() => create();
  static $pb.PbList<Error> createRepeated() => $pb.PbList<Error>();
  @$core.pragma('dart2js:noInline')
  static Error getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Error>(create);
  static Error? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get error => $_getSZ(0);
  @$pb.TagNumber(1)
  set error($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasError() => $_has(0);
  @$pb.TagNumber(1)
  void clearError() => clearField(1);
}

class UpdateNotificationTokenRq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'UpdateNotificationTokenRq', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'global_entry_notify_api'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'userId', protoName: 'userId')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'token')
    ..hasRequiredFields = false
  ;

  UpdateNotificationTokenRq._() : super();
  factory UpdateNotificationTokenRq({
    $core.String? userId,
    $core.String? token,
  }) {
    final _result = create();
    if (userId != null) {
      _result.userId = userId;
    }
    if (token != null) {
      _result.token = token;
    }
    return _result;
  }
  factory UpdateNotificationTokenRq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UpdateNotificationTokenRq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  UpdateNotificationTokenRq clone() => UpdateNotificationTokenRq()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  UpdateNotificationTokenRq copyWith(void Function(UpdateNotificationTokenRq) updates) => super.copyWith((message) => updates(message as UpdateNotificationTokenRq)) as UpdateNotificationTokenRq; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UpdateNotificationTokenRq create() => UpdateNotificationTokenRq._();
  UpdateNotificationTokenRq createEmptyInstance() => create();
  static $pb.PbList<UpdateNotificationTokenRq> createRepeated() => $pb.PbList<UpdateNotificationTokenRq>();
  @$core.pragma('dart2js:noInline')
  static UpdateNotificationTokenRq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UpdateNotificationTokenRq>(create);
  static UpdateNotificationTokenRq? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get userId => $_getSZ(0);
  @$pb.TagNumber(1)
  set userId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasUserId() => $_has(0);
  @$pb.TagNumber(1)
  void clearUserId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get token => $_getSZ(1);
  @$pb.TagNumber(2)
  set token($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasToken() => $_has(1);
  @$pb.TagNumber(2)
  void clearToken() => clearField(2);
}

class UpdateNotificationDetailsRq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'UpdateNotificationDetailsRq', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'global_entry_notify_api'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'userId', protoName: 'userId')
    ..aOM<NotificationDetails>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'notificationDetails', protoName: 'notificationDetails', subBuilder: NotificationDetails.create)
    ..hasRequiredFields = false
  ;

  UpdateNotificationDetailsRq._() : super();
  factory UpdateNotificationDetailsRq({
    $core.String? userId,
    NotificationDetails? notificationDetails,
  }) {
    final _result = create();
    if (userId != null) {
      _result.userId = userId;
    }
    if (notificationDetails != null) {
      _result.notificationDetails = notificationDetails;
    }
    return _result;
  }
  factory UpdateNotificationDetailsRq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UpdateNotificationDetailsRq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  UpdateNotificationDetailsRq clone() => UpdateNotificationDetailsRq()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  UpdateNotificationDetailsRq copyWith(void Function(UpdateNotificationDetailsRq) updates) => super.copyWith((message) => updates(message as UpdateNotificationDetailsRq)) as UpdateNotificationDetailsRq; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UpdateNotificationDetailsRq create() => UpdateNotificationDetailsRq._();
  UpdateNotificationDetailsRq createEmptyInstance() => create();
  static $pb.PbList<UpdateNotificationDetailsRq> createRepeated() => $pb.PbList<UpdateNotificationDetailsRq>();
  @$core.pragma('dart2js:noInline')
  static UpdateNotificationDetailsRq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UpdateNotificationDetailsRq>(create);
  static UpdateNotificationDetailsRq? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get userId => $_getSZ(0);
  @$pb.TagNumber(1)
  set userId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasUserId() => $_has(0);
  @$pb.TagNumber(1)
  void clearUserId() => clearField(1);

  @$pb.TagNumber(2)
  NotificationDetails get notificationDetails => $_getN(1);
  @$pb.TagNumber(2)
  set notificationDetails(NotificationDetails v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasNotificationDetails() => $_has(1);
  @$pb.TagNumber(2)
  void clearNotificationDetails() => clearField(2);
  @$pb.TagNumber(2)
  NotificationDetails ensureNotificationDetails() => $_ensure(1);
}

class DeleteNotificationDetailsRq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'DeleteNotificationDetailsRq', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'global_entry_notify_api'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'userId', protoName: 'userId')
    ..a<$core.int>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'locationId', $pb.PbFieldType.O3, protoName: 'locationId')
    ..hasRequiredFields = false
  ;

  DeleteNotificationDetailsRq._() : super();
  factory DeleteNotificationDetailsRq({
    $core.String? userId,
    $core.int? locationId,
  }) {
    final _result = create();
    if (userId != null) {
      _result.userId = userId;
    }
    if (locationId != null) {
      _result.locationId = locationId;
    }
    return _result;
  }
  factory DeleteNotificationDetailsRq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DeleteNotificationDetailsRq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DeleteNotificationDetailsRq clone() => DeleteNotificationDetailsRq()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DeleteNotificationDetailsRq copyWith(void Function(DeleteNotificationDetailsRq) updates) => super.copyWith((message) => updates(message as DeleteNotificationDetailsRq)) as DeleteNotificationDetailsRq; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static DeleteNotificationDetailsRq create() => DeleteNotificationDetailsRq._();
  DeleteNotificationDetailsRq createEmptyInstance() => create();
  static $pb.PbList<DeleteNotificationDetailsRq> createRepeated() => $pb.PbList<DeleteNotificationDetailsRq>();
  @$core.pragma('dart2js:noInline')
  static DeleteNotificationDetailsRq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DeleteNotificationDetailsRq>(create);
  static DeleteNotificationDetailsRq? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get userId => $_getSZ(0);
  @$pb.TagNumber(1)
  set userId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasUserId() => $_has(0);
  @$pb.TagNumber(1)
  void clearUserId() => clearField(1);

  @$pb.TagNumber(2)
  $core.int get locationId => $_getIZ(1);
  @$pb.TagNumber(2)
  set locationId($core.int v) { $_setSignedInt32(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasLocationId() => $_has(1);
  @$pb.TagNumber(2)
  void clearLocationId() => clearField(2);
}

class NotificationDetails extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'NotificationDetails', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'global_entry_notify_api'), createEmptyInstance: create)
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'locationId', $pb.PbFieldType.O3, protoName: 'locationId')
    ..aOM<$2.Timestamp>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'targetDate', protoName: 'targetDate', subBuilder: $2.Timestamp.create)
    ..aOM<$2.Timestamp>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'appointmentDate', protoName: 'appointmentDate', subBuilder: $2.Timestamp.create)
    ..aOM<$2.Timestamp>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'lastNotifiedDate', protoName: 'lastNotifiedDate', subBuilder: $2.Timestamp.create)
    ..hasRequiredFields = false
  ;

  NotificationDetails._() : super();
  factory NotificationDetails({
    $core.int? locationId,
    $2.Timestamp? targetDate,
    $2.Timestamp? appointmentDate,
    $2.Timestamp? lastNotifiedDate,
  }) {
    final _result = create();
    if (locationId != null) {
      _result.locationId = locationId;
    }
    if (targetDate != null) {
      _result.targetDate = targetDate;
    }
    if (appointmentDate != null) {
      _result.appointmentDate = appointmentDate;
    }
    if (lastNotifiedDate != null) {
      _result.lastNotifiedDate = lastNotifiedDate;
    }
    return _result;
  }
  factory NotificationDetails.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory NotificationDetails.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  NotificationDetails clone() => NotificationDetails()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  NotificationDetails copyWith(void Function(NotificationDetails) updates) => super.copyWith((message) => updates(message as NotificationDetails)) as NotificationDetails; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static NotificationDetails create() => NotificationDetails._();
  NotificationDetails createEmptyInstance() => create();
  static $pb.PbList<NotificationDetails> createRepeated() => $pb.PbList<NotificationDetails>();
  @$core.pragma('dart2js:noInline')
  static NotificationDetails getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<NotificationDetails>(create);
  static NotificationDetails? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get locationId => $_getIZ(0);
  @$pb.TagNumber(1)
  set locationId($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasLocationId() => $_has(0);
  @$pb.TagNumber(1)
  void clearLocationId() => clearField(1);

  @$pb.TagNumber(2)
  $2.Timestamp get targetDate => $_getN(1);
  @$pb.TagNumber(2)
  set targetDate($2.Timestamp v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasTargetDate() => $_has(1);
  @$pb.TagNumber(2)
  void clearTargetDate() => clearField(2);
  @$pb.TagNumber(2)
  $2.Timestamp ensureTargetDate() => $_ensure(1);

  @$pb.TagNumber(3)
  $2.Timestamp get appointmentDate => $_getN(2);
  @$pb.TagNumber(3)
  set appointmentDate($2.Timestamp v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasAppointmentDate() => $_has(2);
  @$pb.TagNumber(3)
  void clearAppointmentDate() => clearField(3);
  @$pb.TagNumber(3)
  $2.Timestamp ensureAppointmentDate() => $_ensure(2);

  @$pb.TagNumber(4)
  $2.Timestamp get lastNotifiedDate => $_getN(3);
  @$pb.TagNumber(4)
  set lastNotifiedDate($2.Timestamp v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasLastNotifiedDate() => $_has(3);
  @$pb.TagNumber(4)
  void clearLastNotifiedDate() => clearField(4);
  @$pb.TagNumber(4)
  $2.Timestamp ensureLastNotifiedDate() => $_ensure(3);
}

class LocationsRp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'LocationsRp', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'global_entry_notify_api'), createEmptyInstance: create)
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'locationId', $pb.PbFieldType.O3, protoName: 'locationId')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'name')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'address')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'addressAdditional', protoName: 'addressAdditional')
    ..aOS(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'city')
    ..aOS(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'state')
    ..aOS(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'postalCode', protoName: 'postalCode')
    ..aOS(8, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'countryCode', protoName: 'countryCode')
    ..aOS(9, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'timezoneData', protoName: 'timezoneData')
    ..aOS(10, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'phoneNumber', protoName: 'phoneNumber')
    ..aOS(11, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'phoneAreaCode', protoName: 'phoneAreaCode')
    ..aOS(12, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'phoneCountryCode', protoName: 'phoneCountryCode')
    ..aOS(13, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Directions', protoName: 'Directions')
    ..aOS(14, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Notes', protoName: 'Notes')
    ..hasRequiredFields = false
  ;

  LocationsRp._() : super();
  factory LocationsRp({
    $core.int? locationId,
    $core.String? name,
    $core.String? address,
    $core.String? addressAdditional,
    $core.String? city,
    $core.String? state,
    $core.String? postalCode,
    $core.String? countryCode,
    $core.String? timezoneData,
    $core.String? phoneNumber,
    $core.String? phoneAreaCode,
    $core.String? phoneCountryCode,
    $core.String? directions,
    $core.String? notes,
  }) {
    final _result = create();
    if (locationId != null) {
      _result.locationId = locationId;
    }
    if (name != null) {
      _result.name = name;
    }
    if (address != null) {
      _result.address = address;
    }
    if (addressAdditional != null) {
      _result.addressAdditional = addressAdditional;
    }
    if (city != null) {
      _result.city = city;
    }
    if (state != null) {
      _result.state = state;
    }
    if (postalCode != null) {
      _result.postalCode = postalCode;
    }
    if (countryCode != null) {
      _result.countryCode = countryCode;
    }
    if (timezoneData != null) {
      _result.timezoneData = timezoneData;
    }
    if (phoneNumber != null) {
      _result.phoneNumber = phoneNumber;
    }
    if (phoneAreaCode != null) {
      _result.phoneAreaCode = phoneAreaCode;
    }
    if (phoneCountryCode != null) {
      _result.phoneCountryCode = phoneCountryCode;
    }
    if (directions != null) {
      _result.directions = directions;
    }
    if (notes != null) {
      _result.notes = notes;
    }
    return _result;
  }
  factory LocationsRp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory LocationsRp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  LocationsRp clone() => LocationsRp()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  LocationsRp copyWith(void Function(LocationsRp) updates) => super.copyWith((message) => updates(message as LocationsRp)) as LocationsRp; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static LocationsRp create() => LocationsRp._();
  LocationsRp createEmptyInstance() => create();
  static $pb.PbList<LocationsRp> createRepeated() => $pb.PbList<LocationsRp>();
  @$core.pragma('dart2js:noInline')
  static LocationsRp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<LocationsRp>(create);
  static LocationsRp? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get locationId => $_getIZ(0);
  @$pb.TagNumber(1)
  set locationId($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasLocationId() => $_has(0);
  @$pb.TagNumber(1)
  void clearLocationId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get name => $_getSZ(1);
  @$pb.TagNumber(2)
  set name($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasName() => $_has(1);
  @$pb.TagNumber(2)
  void clearName() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get address => $_getSZ(2);
  @$pb.TagNumber(3)
  set address($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasAddress() => $_has(2);
  @$pb.TagNumber(3)
  void clearAddress() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get addressAdditional => $_getSZ(3);
  @$pb.TagNumber(4)
  set addressAdditional($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasAddressAdditional() => $_has(3);
  @$pb.TagNumber(4)
  void clearAddressAdditional() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get city => $_getSZ(4);
  @$pb.TagNumber(5)
  set city($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasCity() => $_has(4);
  @$pb.TagNumber(5)
  void clearCity() => clearField(5);

  @$pb.TagNumber(6)
  $core.String get state => $_getSZ(5);
  @$pb.TagNumber(6)
  set state($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasState() => $_has(5);
  @$pb.TagNumber(6)
  void clearState() => clearField(6);

  @$pb.TagNumber(7)
  $core.String get postalCode => $_getSZ(6);
  @$pb.TagNumber(7)
  set postalCode($core.String v) { $_setString(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasPostalCode() => $_has(6);
  @$pb.TagNumber(7)
  void clearPostalCode() => clearField(7);

  @$pb.TagNumber(8)
  $core.String get countryCode => $_getSZ(7);
  @$pb.TagNumber(8)
  set countryCode($core.String v) { $_setString(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasCountryCode() => $_has(7);
  @$pb.TagNumber(8)
  void clearCountryCode() => clearField(8);

  @$pb.TagNumber(9)
  $core.String get timezoneData => $_getSZ(8);
  @$pb.TagNumber(9)
  set timezoneData($core.String v) { $_setString(8, v); }
  @$pb.TagNumber(9)
  $core.bool hasTimezoneData() => $_has(8);
  @$pb.TagNumber(9)
  void clearTimezoneData() => clearField(9);

  @$pb.TagNumber(10)
  $core.String get phoneNumber => $_getSZ(9);
  @$pb.TagNumber(10)
  set phoneNumber($core.String v) { $_setString(9, v); }
  @$pb.TagNumber(10)
  $core.bool hasPhoneNumber() => $_has(9);
  @$pb.TagNumber(10)
  void clearPhoneNumber() => clearField(10);

  @$pb.TagNumber(11)
  $core.String get phoneAreaCode => $_getSZ(10);
  @$pb.TagNumber(11)
  set phoneAreaCode($core.String v) { $_setString(10, v); }
  @$pb.TagNumber(11)
  $core.bool hasPhoneAreaCode() => $_has(10);
  @$pb.TagNumber(11)
  void clearPhoneAreaCode() => clearField(11);

  @$pb.TagNumber(12)
  $core.String get phoneCountryCode => $_getSZ(11);
  @$pb.TagNumber(12)
  set phoneCountryCode($core.String v) { $_setString(11, v); }
  @$pb.TagNumber(12)
  $core.bool hasPhoneCountryCode() => $_has(11);
  @$pb.TagNumber(12)
  void clearPhoneCountryCode() => clearField(12);

  @$pb.TagNumber(13)
  $core.String get directions => $_getSZ(12);
  @$pb.TagNumber(13)
  set directions($core.String v) { $_setString(12, v); }
  @$pb.TagNumber(13)
  $core.bool hasDirections() => $_has(12);
  @$pb.TagNumber(13)
  void clearDirections() => clearField(13);

  @$pb.TagNumber(14)
  $core.String get notes => $_getSZ(13);
  @$pb.TagNumber(14)
  set notes($core.String v) { $_setString(13, v); }
  @$pb.TagNumber(14)
  $core.bool hasNotes() => $_has(13);
  @$pb.TagNumber(14)
  void clearNotes() => clearField(14);
}

class GetNotificationDetailsRq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'GetNotificationDetailsRq', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'global_entry_notify_api'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'userId', protoName: 'userId')
    ..hasRequiredFields = false
  ;

  GetNotificationDetailsRq._() : super();
  factory GetNotificationDetailsRq({
    $core.String? userId,
  }) {
    final _result = create();
    if (userId != null) {
      _result.userId = userId;
    }
    return _result;
  }
  factory GetNotificationDetailsRq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetNotificationDetailsRq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetNotificationDetailsRq clone() => GetNotificationDetailsRq()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetNotificationDetailsRq copyWith(void Function(GetNotificationDetailsRq) updates) => super.copyWith((message) => updates(message as GetNotificationDetailsRq)) as GetNotificationDetailsRq; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GetNotificationDetailsRq create() => GetNotificationDetailsRq._();
  GetNotificationDetailsRq createEmptyInstance() => create();
  static $pb.PbList<GetNotificationDetailsRq> createRepeated() => $pb.PbList<GetNotificationDetailsRq>();
  @$core.pragma('dart2js:noInline')
  static GetNotificationDetailsRq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetNotificationDetailsRq>(create);
  static GetNotificationDetailsRq? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get userId => $_getSZ(0);
  @$pb.TagNumber(1)
  set userId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasUserId() => $_has(0);
  @$pb.TagNumber(1)
  void clearUserId() => clearField(1);
}

class GetNotificationDetailsRp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'GetNotificationDetailsRp', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'global_entry_notify_api'), createEmptyInstance: create)
    ..pc<NotificationDetails>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'notificationDetails', $pb.PbFieldType.PM, protoName: 'notificationDetails', subBuilder: NotificationDetails.create)
    ..hasRequiredFields = false
  ;

  GetNotificationDetailsRp._() : super();
  factory GetNotificationDetailsRp({
    $core.Iterable<NotificationDetails>? notificationDetails,
  }) {
    final _result = create();
    if (notificationDetails != null) {
      _result.notificationDetails.addAll(notificationDetails);
    }
    return _result;
  }
  factory GetNotificationDetailsRp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetNotificationDetailsRp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetNotificationDetailsRp clone() => GetNotificationDetailsRp()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetNotificationDetailsRp copyWith(void Function(GetNotificationDetailsRp) updates) => super.copyWith((message) => updates(message as GetNotificationDetailsRp)) as GetNotificationDetailsRp; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GetNotificationDetailsRp create() => GetNotificationDetailsRp._();
  GetNotificationDetailsRp createEmptyInstance() => create();
  static $pb.PbList<GetNotificationDetailsRp> createRepeated() => $pb.PbList<GetNotificationDetailsRp>();
  @$core.pragma('dart2js:noInline')
  static GetNotificationDetailsRp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetNotificationDetailsRp>(create);
  static GetNotificationDetailsRp? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<NotificationDetails> get notificationDetails => $_getList(0);
}

