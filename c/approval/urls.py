from django.urls import path
from approval.apis import approval

urlpatterns = [
    path('v1/subscribe', approval.FApprovalView.as_view(), name='subscribe'),
    path('v1/put-deploy-process',
         approval.SendDeployResultView.as_view(),
         name='put-deploy-process'),
]