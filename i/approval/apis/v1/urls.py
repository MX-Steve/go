from django.urls import path
from approval.apis.v1.approval_bind import ApprovalBindView
from approval.apis.v1.instance import FInstanceView, FInstanceListView, ChoicesView
from approval.apis.v1.event_listen import ListenApprovalEventView
from approval.apis.v1.send_result import SendDeployResultView, SendResultToGroup, SendResultToGroup2

urlpatterns = [
    path('subscribe', ApprovalBindView.as_view(), name='subscribe'),
    path('instance', FInstanceView.as_view(), name="instance"),
    path('instances', FInstanceListView.as_view(), name="instances"),
    path('event', ListenApprovalEventView.as_view(), name='event'),
    path('send-result', SendDeployResultView.as_view(), name='send-result'),
    path('send-result-to-group', SendResultToGroup.as_view(), name='send-result-to-group'),
    path('send-result-to-group2', SendResultToGroup2.as_view(), name='send-result-to-group2'),
    path('choices', ChoicesView.as_view(), name='choices'),
]