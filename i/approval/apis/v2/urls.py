from django.urls import path
from approval.apis.v2.approval import FApprovalView, FAFormView, SubscribeApprovalView, FANodeView, ChangeVersionView
from approval.apis.v2.instance import FIFormView, FInstanceValueView, FITaskView
from approval.apis.v2.choices import ChoicesView
from approval.apis.v2.send_result import SendDeployResultView, SendResultToSelf2, SendResultToGroup2, SendResultToSelf
from approval.apis.v2.instance_process import ProcessView, InfosView, DeployDirectlyView
from approval.apis.v2.dashboard import DashboardView
from approval.apis.v2.deploy_history import DeployHistoryView
from approval.apis.v2.uploads import UploadFilesView, DownloadFilesView
from approval.apis.v2.task import FITaskUserView
from approval.views import grpc_test, kafka_get

urlpatterns = [
    path('approval', FApprovalView.as_view(), name='approval'),
    path('form', FAFormView.as_view(), name='form'),
    path('subscribe', SubscribeApprovalView.as_view(), name='subscribe'),
    path('change-version', ChangeVersionView.as_view(), name='change-version'),
    path('node', FANodeView.as_view(), name='node'),
    path('instance', FInstanceValueView.as_view(), name='instance'),
    path('fiform', FIFormView.as_view(), name='fiform'),
    path('tasks', FITaskView.as_view(), name='tasks'),
    path('task-modify-user', FITaskUserView.as_view(), name='task-modify-user'),
    path('send-result', SendDeployResultView.as_view(), name='send-result'),
    path('send-result-to-group2', SendResultToGroup2.as_view(), name='send-result-to-group2'),
    path('send-result-test', SendResultToSelf.as_view(), name='send-result-test'),
    path('send-result-self',
         SendResultToSelf2.as_view(),
         name='send-result-self'),
    path('choices', ChoicesView.as_view(), name='choices'),
    path('process', ProcessView.as_view(), name='process'),
    path('direct-deploy', DeployDirectlyView.as_view(), name='direct-deploy'),
    path('infos', InfosView.as_view(), name='infos'),
    path("summary-approval", DashboardView.as_view(), name="summary-approval" ),
    path("deploy-history", DeployHistoryView.as_view(), name="deploy-history"),
    path('upload-file', UploadFilesView.as_view(), name="upload-file"),
    path('download-file', DownloadFilesView.as_view(), name="download-file"),
    path('grpc-test', grpc_test, name='grpc-test'),
    path('kafka-get',kafka_get, name='kafka-get'),
]