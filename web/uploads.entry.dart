import 'dart:convert';
import 'dart:html';

import 'package:mango_ui/services/uploadapi.dart';

void main() {
  document.body.onClick.matches('.deleter').listen(onDeleteClick);
}

void onDeleteClick(MouseEvent e) async {
  final targt = e.matchingTarget;

  if (targt is ButtonElement) {
    final toDelete = targt.dataset["key"];
    final warn = "Are you sure you want to Delete ${toDelete}?";
    if (window.confirm(warn)) {
      final req = await removeUpload(toDelete);
      final resp = jsonDecode(req.response);

      if (req.status == 200) {
        window.alert(resp["Data"]);
      } else {
        print(resp["Error"]);
      }
    }
  }
}
