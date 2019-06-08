import 'package:Admin.APP/keys.dart';
import 'package:Admin.APP/rolesform.dart';

void main() {
  print('Running User View');

  new RolesForm("#frmRoleCreate", getObjKey(), "#btnSave", "#btnAddRole", "#tblRoles");
}