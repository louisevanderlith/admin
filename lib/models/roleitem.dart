import 'dart:html';

import 'package:mango_ui/bodies/role.dart';

//RoleItem is used within tables to display data in a row.
class RoleItem {
  LabelElement _application;
  ElementList<RadioButtonInputElement> _roletypes;
  
  bool _loaded;

  RoleItem(String applicationElem, String roletypeElem) {
    _application = querySelector(applicationElem);
    _roletypes = querySelectorAll(roletypeElem);
    _loaded = _application != null && _roletypes.length != 0;
  }
  
  String get application {
    return _application.text;
  }

  void set application(String appName) {
    _application.text = appName;
  }

  num get roletype {
    for (var i = 0; i < _roletypes.length; i++) {
      final curr = _roletypes[i];
      if (curr.checked) {
        return num.parse(curr.value);
      }
    }

    return 3;
  }

  void set roletype(num role) {
    for (var i = 0; i < _roletypes.length; i++) {
      final curr = _roletypes[i];
      curr.checked = curr.value == role;
    }
  }

  bool loaded() {
    return _loaded;
  }

  Role toDTO() {
    return new Role(application, roletype);
  }
}