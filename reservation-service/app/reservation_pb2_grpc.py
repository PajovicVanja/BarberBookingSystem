# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from . import reservation_pb2 as reservation__pb2

GRPC_GENERATED_VERSION = '1.71.0'
GRPC_VERSION = grpc.__version__
_version_not_supported = False

try:
    from grpc._utilities import first_version_is_lower
    _version_not_supported = first_version_is_lower(GRPC_VERSION, GRPC_GENERATED_VERSION)
except ImportError:
    _version_not_supported = True

if _version_not_supported:
    raise RuntimeError(
        f'The grpc package installed is at version {GRPC_VERSION},'
        + f' but the generated code in reservation_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
    )


class ReservationServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreateReservation = channel.unary_unary(
                '/reservation.ReservationService/CreateReservation',
                request_serializer=reservation__pb2.ReservationRequest.SerializeToString,
                response_deserializer=reservation__pb2.ReservationResponse.FromString,
                _registered_method=True)
        self.GetReservation = channel.unary_unary(
                '/reservation.ReservationService/GetReservation',
                request_serializer=reservation__pb2.GetReservationRequest.SerializeToString,
                response_deserializer=reservation__pb2.ReservationResponse.FromString,
                _registered_method=True)
        self.CancelReservation = channel.unary_unary(
                '/reservation.ReservationService/CancelReservation',
                request_serializer=reservation__pb2.CancelReservationRequest.SerializeToString,
                response_deserializer=reservation__pb2.ReservationResponse.FromString,
                _registered_method=True)
        self.UpdateReservation = channel.unary_unary(
                '/reservation.ReservationService/UpdateReservation',
                request_serializer=reservation__pb2.UpdateReservationRequest.SerializeToString,
                response_deserializer=reservation__pb2.ReservationResponse.FromString,
                _registered_method=True)
        self.ListUserReservations = channel.unary_unary(
                '/reservation.ReservationService/ListUserReservations',
                request_serializer=reservation__pb2.ListUserReservationsRequest.SerializeToString,
                response_deserializer=reservation__pb2.ReservationListResponse.FromString,
                _registered_method=True)
        self.ConfirmPayment = channel.unary_unary(
                '/reservation.ReservationService/ConfirmPayment',
                request_serializer=reservation__pb2.ConfirmPaymentRequest.SerializeToString,
                response_deserializer=reservation__pb2.ReservationResponse.FromString,
                _registered_method=True)


class ReservationServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def CreateReservation(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetReservation(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CancelReservation(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateReservation(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListUserReservations(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ConfirmPayment(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ReservationServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CreateReservation': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateReservation,
                    request_deserializer=reservation__pb2.ReservationRequest.FromString,
                    response_serializer=reservation__pb2.ReservationResponse.SerializeToString,
            ),
            'GetReservation': grpc.unary_unary_rpc_method_handler(
                    servicer.GetReservation,
                    request_deserializer=reservation__pb2.GetReservationRequest.FromString,
                    response_serializer=reservation__pb2.ReservationResponse.SerializeToString,
            ),
            'CancelReservation': grpc.unary_unary_rpc_method_handler(
                    servicer.CancelReservation,
                    request_deserializer=reservation__pb2.CancelReservationRequest.FromString,
                    response_serializer=reservation__pb2.ReservationResponse.SerializeToString,
            ),
            'UpdateReservation': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateReservation,
                    request_deserializer=reservation__pb2.UpdateReservationRequest.FromString,
                    response_serializer=reservation__pb2.ReservationResponse.SerializeToString,
            ),
            'ListUserReservations': grpc.unary_unary_rpc_method_handler(
                    servicer.ListUserReservations,
                    request_deserializer=reservation__pb2.ListUserReservationsRequest.FromString,
                    response_serializer=reservation__pb2.ReservationListResponse.SerializeToString,
            ),
            'ConfirmPayment': grpc.unary_unary_rpc_method_handler(
                    servicer.ConfirmPayment,
                    request_deserializer=reservation__pb2.ConfirmPaymentRequest.FromString,
                    response_serializer=reservation__pb2.ReservationResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'reservation.ReservationService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('reservation.ReservationService', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class ReservationService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def CreateReservation(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/reservation.ReservationService/CreateReservation',
            reservation__pb2.ReservationRequest.SerializeToString,
            reservation__pb2.ReservationResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def GetReservation(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/reservation.ReservationService/GetReservation',
            reservation__pb2.GetReservationRequest.SerializeToString,
            reservation__pb2.ReservationResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def CancelReservation(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/reservation.ReservationService/CancelReservation',
            reservation__pb2.CancelReservationRequest.SerializeToString,
            reservation__pb2.ReservationResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def UpdateReservation(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/reservation.ReservationService/UpdateReservation',
            reservation__pb2.UpdateReservationRequest.SerializeToString,
            reservation__pb2.ReservationResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def ListUserReservations(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/reservation.ReservationService/ListUserReservations',
            reservation__pb2.ListUserReservationsRequest.SerializeToString,
            reservation__pb2.ReservationListResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def ConfirmPayment(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/reservation.ReservationService/ConfirmPayment',
            reservation__pb2.ConfirmPaymentRequest.SerializeToString,
            reservation__pb2.ReservationResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
