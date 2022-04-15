from django.db import models

class Host(models.Model):
    ip_address = models.CharField(max_length=100)
    username = models.CharField(max_length=100)
    port = models.IntegerField(null=True)
    pkey = models.TextField(null=True)