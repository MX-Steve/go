from approval.models import User, FANode


def putNodes(f_approval_id, node_list):
    FANode.objects.filter(f_approval_id=f_approval_id).update(del_tag=1)
    user = User.objects.filter(username="admin").first()
    user_id = user.id
    nodes = [{
        "name": "发起",
        "need_approver": False,
        "need_cc_user": False,
        "node_id": "b078ffd28db767c502ac367053f6e0ac",
        "node_type": "AND"
    }]
    for node in node_list:
        if node["name"] not in ["发起", "结束"]:
            nodes.append(node)
    nodes.append({
        "name": "结束",
        "need_approver": False,
        "need_cc_user": False,
        "node_id": "b1a326c06d88bf042f73d70f50197905",
        "node_type": "AND"
    })
    i = 1
    for item in nodes:
        if item["name"] in ["发起", "结束"]:
            approval_type = "AUTO_PASS"
        else:
            approval_type = "MANUAL"
        print("name: ", item["name"])
        print("need_approver: ", item["need_approver"])
        print("node_type: ", item["node_type"])
        print("user_id: ", user_id)
        print("approval_type: ", approval_type)
        print("serial_number: ", i)
        print("f_approval_id: ", f_approval_id)
        print("del_tag: ", 0)
        print("-----")
        FANode.objects.create(name=item["name"],
                              need_approver=item["need_approver"],
                              node_type=item["node_type"],
                              user_id=user_id,
                              approval_type=approval_type,
                              serial_number=i,
                              f_approval_id=f_approval_id,
                              del_tag=0)
        i += 1


if __name__ == '__main__':
    node_list = [{
        "name": "操作人",
        "need_approver": False,
        "need_cc_user": False,
        "node_id": "38d5d522a17022e7410bae6c09c11749",
        "node_type": "AND"
    }, {
        "name": "项目经理审批",
        "need_approver": True,
        "need_cc_user": False,
        "node_id": "3c410e0961df204739b06de1ce4873ec",
        "node_type": "AND"
    }, {
        "name": "审批",
        "need_approver": False,
        "need_cc_user": False,
        "node_id": "50d41b165717195c07babd8ddbdd1f2e",
        "node_type": "AND"
    }, {
        "name": "结束",
        "need_approver": False,
        "need_cc_user": False,
        "node_id": "b1a326c06d88bf042f73d70f50197905",
        "node_type": "AND"
    }, {
        "name": "发起",
        "need_approver": False,
        "need_cc_user": False,
        "node_id": "b078ffd28db767c502ac367053f6e0ac",
        "node_type": "AND"
    }]
    f_approval_id = 1
    putNodes(f_approval_id, node_list)
