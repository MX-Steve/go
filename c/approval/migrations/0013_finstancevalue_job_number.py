# Generated by Django 3.2.7 on 2022-01-27 03:20

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('approval', '0012_auto_20220127_0942'),
    ]

    operations = [
        migrations.AddField(
            model_name='finstancevalue',
            name='job_number',
            field=models.IntegerField(default=0, help_text='job 编号'),
        ),
    ]
