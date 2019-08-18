import 'dart:convert';
import 'dart:html';

import 'package:Admin.APP/models/createroleitem.dart';
import 'package:mango_ui/bodies/key.dart';
import 'package:mango_ui/bodies/role.dart';
import 'package:mango_ui/formstate.dart';
import 'package:mango_ui/services/secureapi.dart';

import 'models/roleitem.dart';

class RolesForm extends FormState {
  Element _tblRoles;
  List<Role> _items;
  Key _objKey;

  RolesForm(String idElem, Key objKey, String btnSubmit, String btnAdd,
      String tblRoles)
      : super(idElem, btnSubmit) {
    _items = findRoles();
    _tblRoles = querySelector(tblRoles);
    _objKey = objKey;

    querySelector(btnAdd).onClick.listen(onAddClick);
    querySelector(btnSubmit).onClick.listen(onSubmitClick);
  }

  void onAddClick(MouseEvent e) async {
    final roleItem = new CreateRoleItem();
    await roleItem.display();

    addRole(roleItem);
  }

  void onSubmitClick(MouseEvent e) async {
    if (isFormValid()) {
      disableSubmit(true);
      final req = await updateRoles(_objKey, items);
      var result = jsonDecode(req.response);
      window.alert(result);
    }
  }

  List<Role> get items {
    return _items;
  }

  List<Role> findRoles() {
    var hasItem = false;
    var result = new List<Role>();
    var indx = 0;

    do {
      var item =
          new RoleItem('#lblAppName${indx}', 'input[name=answer${indx}]');
      hasItem = item.loaded();

      if (hasItem) {
        result.add(item.role);
      }

      indx++;
    } while (hasItem);

    return result;
  }

  void addRole(CreateRoleItem obj) {
    //add elements
    var indx = items.length;

    _tblRoles.children.add(obj.toHtml(indx));

    var item = new RoleItem('#lblAppName${indx}', 'input[name=answer${indx}]');
    _items.add(item.role);
  }
}
