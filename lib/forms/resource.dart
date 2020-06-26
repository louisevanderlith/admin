import 'dart:convert';
import 'dart:html';

import 'package:mango_secure/bodies/resource.dart';
import 'package:mango_secure/resourceapi.dart';
import 'package:mango_ui/formstate.dart';
import 'package:mango_ui/keys.dart';

class ResourceForm extends FormState {
  Key _objKey;
  TextInputElement txtName;
  TextInputElement txtDisplayName;
  PasswordInputElement txtSecret;
  TextInputElement txtNewNeed;
  UListElement lstNeeds;

  ResourceForm(Key objKey) : super("#frmResource", "#btnSave") {
    _objKey = objKey;

    txtName = querySelector("#txtName");
    txtDisplayName = querySelector("#txtDisplayName");
    txtSecret = querySelector("#txtSecret");
    txtNewNeed = querySelector("#txtNewNeed");
    lstNeeds = querySelector("#lstNeeds");

    querySelector("#btnSave").onClick.listen(onSubmitClick);
    querySelector("#btnAddNeed").onClick.listen(onNeedAddClick);
  }

  String get name {
    return txtName.text;
  }

  String get displayname {
    return txtDisplayName.text;
  }

  String get secret {
    return txtSecret.text;
  }

  List<String> get needs {
    return lstNeeds.children.map((e) => e.text).toList();
  }

  void onNeedAddClick(MouseEvent e) {
    final item = new Element.li();
    item.text = txtNewNeed.text;

    lstNeeds.children.add(item);
    txtNewNeed.text = "";
  }

  void onSubmitClick(MouseEvent e) async {
    if (isFormValid()) {
      disableSubmit(true);

      final obj =
          new Resource(this.name, this.displayname, this.secret, this.needs);

      HttpRequest req;
      if (_objKey.toJson() != "0`0") {
        req = await updateResource(_objKey, obj);
      } else {
        req = await createResource(obj);
      }

      var result = jsonDecode(req.response);

      if (req.status == 200) {
        final data = result['Data'];
        final rec = data['Record'];

        if (rec != null) {
          final key = rec['K'];

          _objKey = key;
        }
      }
    }
  }
}
