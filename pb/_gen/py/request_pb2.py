# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: request.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='request.proto',
  package='',
  syntax='proto3',
  serialized_options=_b('\n\017com.reworld.wwwB\010TeamworkZ\004.;pb\252\002\020Reworld.Teamwork'),
  serialized_pb=_b('\n\rrequest.proto\"\x85\x01\n\x07Request\x12\x1b\n\x06\x41\x63tion\x18\x01 \x01(\x0e\x32\x0b.PingAction\x12\x0e\n\x06TeamId\x18\x02 \x01(\x03\x12\x0c\n\x04UUID\x18\x03 \x01(\t\x12\x11\n\tTimestamp\x18\x04 \x01(\x03\x12\x0c\n\x04Sign\x18\x05 \x01(\t\x12\x10\n\x08receiver\x18\x06 \x03(\t\x12\x0c\n\x04\x44\x61ta\x18\x07 \x01(\x0c*K\n\nPingAction\x12\r\n\tHeartbeat\x10\x00\x12\x0e\n\nReqReplica\x10\x01\x12\x11\n\rReqDistribute\x10\x02\x12\x0b\n\x07Offline\x10\x03\x42\x34\n\x0f\x63om.reworld.wwwB\x08TeamworkZ\x04.;pb\xaa\x02\x10Reworld.Teamworkb\x06proto3')
)

_PINGACTION = _descriptor.EnumDescriptor(
  name='PingAction',
  full_name='PingAction',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='Heartbeat', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ReqReplica', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ReqDistribute', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Offline', index=3, number=3,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=153,
  serialized_end=228,
)
_sym_db.RegisterEnumDescriptor(_PINGACTION)

PingAction = enum_type_wrapper.EnumTypeWrapper(_PINGACTION)
Heartbeat = 0
ReqReplica = 1
ReqDistribute = 2
Offline = 3



_REQUEST = _descriptor.Descriptor(
  name='Request',
  full_name='Request',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='Action', full_name='Request.Action', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='TeamId', full_name='Request.TeamId', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='UUID', full_name='Request.UUID', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='Timestamp', full_name='Request.Timestamp', index=3,
      number=4, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='Sign', full_name='Request.Sign', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='receiver', full_name='Request.receiver', index=5,
      number=6, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='Data', full_name='Request.Data', index=6,
      number=7, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=18,
  serialized_end=151,
)

_REQUEST.fields_by_name['Action'].enum_type = _PINGACTION
DESCRIPTOR.message_types_by_name['Request'] = _REQUEST
DESCRIPTOR.enum_types_by_name['PingAction'] = _PINGACTION
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Request = _reflection.GeneratedProtocolMessageType('Request', (_message.Message,), {
  'DESCRIPTOR' : _REQUEST,
  '__module__' : 'request_pb2'
  # @@protoc_insertion_point(class_scope:Request)
  })
_sym_db.RegisterMessage(Request)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
