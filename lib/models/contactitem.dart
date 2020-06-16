import 'dart:html';

import 'package:mango_secure/bodies/contact.dart';

class ContactItem {
  TextInputElement txtIcon;
  TextInputElement txtName;
  TextInputElement txtValue;
  bool _loaded;

  ContactItem(String iconId, String nameId, String valueId) {
    txtIcon = querySelector(iconId);
    txtName = querySelector(nameId);
    txtValue = querySelector(valueId);

    _loaded = txtIcon != null && txtName != null && txtValue != null;
  }

  String get icon {
    return txtIcon.value;
  }

  String get name {
    return txtName.value;
  }

  String get value {
    return txtValue.value;
  }

  bool loaded() {
    return _loaded;
  }

  Contact toDTO() {
    return new Contact(icon, name, value);
  }
}
