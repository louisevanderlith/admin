import 'dart:convert';
import 'dart:html';

import 'package:Admin.APP/services/blogapi.dart';

void main() {
  querySelector('#btnAdd').onClick.listen(onAddClick);
  querySelectorAll('.deleter').onClick.listen(onDeleteClick);
}

void onAddClick(MouseEvent e) async {
  final result = await createArticle('New Article', 'System');
  var obj = jsonDecode(result.response);

  if (result.status == 200) {
    final data = obj['Data'];
    final rec = data['Record'];
    final key = rec['K'];

    final redir = "/blog/${key}";
    window.location.replace(redir);
  } else {
    print(obj['Error']);
  }
}

void onDeleteClick(MouseEvent e) async {
  final targt = e.target;
  
  if (targt is ButtonElement) {
    final toDelete = targt.dataset["key"];
    final warn = "Are you sure you want to Delete ${toDelete}?";
    if (window.confirm(warn)) {
      final req = await removeArticle(toDelete);
      final resp = jsonDecode(req.response);

      if (req.status == 200) {
        window.alert(resp["Data"]);
      } else {
        print(resp["Error"]);
      }
    }
  }
}
