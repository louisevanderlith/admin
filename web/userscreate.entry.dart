import 'package:mango_ui/keys.dart';
import 'package:Admin.APP/rolesform.dart';

void main() {
  print('Running User View');

  new RolesForm(
      "#frmRoleCreate", getObjKey(), "#btnSubmit", "#btnAddRole", "#tblRoles");
}
