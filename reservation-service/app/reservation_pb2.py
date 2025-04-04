# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: reservation.proto
# Protobuf Python Version: 5.29.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    29,
    0,
    '',
    'reservation.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x11reservation.proto\x12\x0breservation\"J\n\x12ReservationRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x11\n\tbarber_id\x18\x02 \x01(\t\x12\x10\n\x08\x64\x61tetime\x18\x03 \x01(\t\"#\n\x15GetReservationRequest\x12\n\n\x02id\x18\x01 \x01(\t\"&\n\x18\x43\x61ncelReservationRequest\x12\n\n\x02id\x18\x01 \x01(\t\"8\n\x18UpdateReservationRequest\x12\n\n\x02id\x18\x01 \x01(\t\x12\x10\n\x08\x64\x61tetime\x18\x02 \x01(\t\".\n\x1bListUserReservationsRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\"#\n\x15\x43onfirmPaymentRequest\x12\n\n\x02id\x18\x01 \x01(\t\"_\n\x0bReservation\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0f\n\x07user_id\x18\x02 \x01(\t\x12\x11\n\tbarber_id\x18\x03 \x01(\t\x12\x10\n\x08\x64\x61tetime\x18\x04 \x01(\t\x12\x0e\n\x06status\x18\x05 \x01(\t\"f\n\x13ReservationResponse\x12\x0f\n\x07success\x18\x01 \x01(\x08\x12\x0f\n\x07message\x18\x02 \x01(\t\x12-\n\x0breservation\x18\x03 \x01(\x0b\x32\x18.reservation.Reservation\"I\n\x17ReservationListResponse\x12.\n\x0creservations\x18\x01 \x03(\x0b\x32\x18.reservation.Reservation2\xc0\x04\n\x12ReservationService\x12V\n\x11\x43reateReservation\x12\x1f.reservation.ReservationRequest\x1a .reservation.ReservationResponse\x12V\n\x0eGetReservation\x12\".reservation.GetReservationRequest\x1a .reservation.ReservationResponse\x12\\\n\x11\x43\x61ncelReservation\x12%.reservation.CancelReservationRequest\x1a .reservation.ReservationResponse\x12\\\n\x11UpdateReservation\x12%.reservation.UpdateReservationRequest\x1a .reservation.ReservationResponse\x12\x66\n\x14ListUserReservations\x12(.reservation.ListUserReservationsRequest\x1a$.reservation.ReservationListResponse\x12V\n\x0e\x43onfirmPayment\x12\".reservation.ConfirmPaymentRequest\x1a .reservation.ReservationResponseb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'reservation_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  DESCRIPTOR._loaded_options = None
  _globals['_RESERVATIONREQUEST']._serialized_start=34
  _globals['_RESERVATIONREQUEST']._serialized_end=108
  _globals['_GETRESERVATIONREQUEST']._serialized_start=110
  _globals['_GETRESERVATIONREQUEST']._serialized_end=145
  _globals['_CANCELRESERVATIONREQUEST']._serialized_start=147
  _globals['_CANCELRESERVATIONREQUEST']._serialized_end=185
  _globals['_UPDATERESERVATIONREQUEST']._serialized_start=187
  _globals['_UPDATERESERVATIONREQUEST']._serialized_end=243
  _globals['_LISTUSERRESERVATIONSREQUEST']._serialized_start=245
  _globals['_LISTUSERRESERVATIONSREQUEST']._serialized_end=291
  _globals['_CONFIRMPAYMENTREQUEST']._serialized_start=293
  _globals['_CONFIRMPAYMENTREQUEST']._serialized_end=328
  _globals['_RESERVATION']._serialized_start=330
  _globals['_RESERVATION']._serialized_end=425
  _globals['_RESERVATIONRESPONSE']._serialized_start=427
  _globals['_RESERVATIONRESPONSE']._serialized_end=529
  _globals['_RESERVATIONLISTRESPONSE']._serialized_start=531
  _globals['_RESERVATIONLISTRESPONSE']._serialized_end=604
  _globals['_RESERVATIONSERVICE']._serialized_start=607
  _globals['_RESERVATIONSERVICE']._serialized_end=1183
# @@protoc_insertion_point(module_scope)
