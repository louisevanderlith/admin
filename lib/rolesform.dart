import 'dart:convert';
import 'dart:html';

import 'package:Admin.APP/models/createroleitem.dart';
import 'package:Admin.APP/services/secureapi.dart';

import 'formstate.dart';
import 'models/roleitem.dart';

class RolesForm extends FormState {
  Element _tblRoles;
  List<RoleItem> _items;
  String _objKey;

  RolesForm(String idElem, String objKey, String btnSubmit, String btnAdd, String tblRoles)
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

  List<RoleItem> get items {
    return _items;
  }

  List<RoleItem> findRoles() {
    var hasItem = false;
    var result = new List<RoleItem>();
    var indx = 0;

    do {
      var item =
          new RoleItem('#lblAppName${indx}', 'input[name=answer${indx}]');
      hasItem = item.loaded();

      if (hasItem) {
        result.add(item);
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
    _items.add(item);
  }
}
