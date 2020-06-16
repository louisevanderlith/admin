import 'dart:html';

import 'package:mango_ui/formstate.dart';
import 'package:mango_ui/keys.dart';

class ProfileForm extends FormState {
  Key _objKey;
  TextInputElement txtTitle;
  TextAreaElement txtDescription;
  ContactsForm frmContacts;
  EndpointsForm frmEndpoints;
  CodesForm frmCodes;
  UListElement lstTerms;

  String get title {
    return txtTitle.value;
  }

  String get description {
  return txtDescription.value;
  }

  List<Contact> get contacts {
    return frmContacts.items;
  }

  void onSubmitClick(MouseEvent e) async {
    if (isFormValid()) {
      disableSubmit(true);
      final obj = new Profile(
          title, intro, category, imageKey, content, writtenby, public).;

      HttpRequest req;
      if (_objKey.toJson() != "0`0") {
        req = await updateArticle(_objKey, obj);
      } else {
        req = await createArticle(obj);
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