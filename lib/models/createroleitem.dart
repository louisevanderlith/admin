import 'dart:async';
import 'dart:html';

import 'package:Admin.APP/trustvalidator.dart';

class CreateRoleItem {
  Completer _compltr;
  SelectElement _application;
  ElementList<RadioButtonInputElement> _roletypes;
  DivElement _modal;

  bool _loaded;

  CreateRoleItem() {
    _application = querySelector('#cboApplicationAdd');
    _roletypes = querySelectorAll('input[name=answerAdd]');
    _modal = querySelector('#theModal');

    _loaded = _application != null && _roletypes != null;

    final _btnSubmit = querySelector('#btnCreate');
    _compltr = new Completer<Object>();
    _btnSubmit.onClick.listen(submitItem);
  }

  String get application {
    return _application.value;
  }

  String get roletype {
    for (var i = 0; i < _roletypes.length; i++) {
      final curr = _roletypes[i];
      if (curr.checked) {
        return curr.value;
      }
    }

    return "";
  }

  bool loaded() {
    return _loaded;
  }

  Future<Object> display() {
    _modal.classes.add('is-active');

    return _compltr.future;
  }

  void submitItem(MouseEvent e) {
    e.preventDefault();
    _modal.classes.remove('is-active');
    _compltr.complete();
  }

  Element toHtml(num index) {
    var radios = '';
    for (var i = 0; i < 3; i++) {
      final checked = roletype == i.toString() ? 'checked' : '';
      var rad = '''<td><div class="control">
            <label class="radio">
                <input type="radio" ${checked} value="${i}"
                    name="answer${index}"/>
            </label>
        </div></td>''';

      radios += rad;
    }

    var schema = '''
      <td> 
        <label id="lblAppName${index}" class="label" >${application}</label>
      </td>
      ${radios}
        ''';

    var row = Element.tr();
    row.setInnerHtml(schema, validator: new TrustedNodeValidator());
    return row;
  }
}
