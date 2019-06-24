import 'dart:convert';
import 'dart:html';

import 'package:Admin.APP/formstate.dart';
import 'package:Admin.APP/services/blogapi.dart';
import 'package:Admin.APP/services/uploadapi.dart';

class ArticleForm extends FormState {
  String _objKey;
  TextInputElement _title;
  TextInputElement _intro;
  SelectElement _categories;
  DivElement _content;
  FileUploadInputElement _headImage;

  ArticleForm(String idElem, String objKey, String titleElem, String introElem, String categoriesElem, String contentElem,
      String imageElem, String previewBtn, String publishBtn, String submitBtn)
      : super(idElem, submitBtn) {
    _objKey = objKey;
    _title = querySelector(titleElem);
    _intro = querySelector(introElem);
    _categories = querySelector(categoriesElem);
    _content = querySelector(contentElem);
    _headImage = querySelector(imageElem);

    querySelector(submitBtn).onClick.listen(onSubmitClick);
    querySelector(previewBtn).onClick.listen(onPreviewClick);
    querySelector(publishBtn).onClick.listen(onPublishClick);

    _headImage.onChange.listen(uploadFile);

    //Editor events
    querySelectorAll('#editCtrls')
        .onClick
        .matches('span.button')
        .listen(onEditorCtrlClick);
  }

  String get title {
    return _title.value;
  }

  String get intro {
    return _intro.value;
  }

  String get category {
    return _categories.value;
  }
  
  String get content {
    return _content.innerHtml;
  }

  String get imageKey {
    return _headImage.dataset['id'];
  }

  void onSubmitClick(MouseEvent e) async {
    if (isFormValid()) {
      disableSubmit(true);
      final req =
          await updateArticle(_objKey, title, intro, category, content, imageKey, 'System');
      var result = jsonDecode(req.response);

      print(result);
    }
  }

  void onPreviewClick(MouseEvent e) async {
    if (isFormValid()) {
      disableSubmit(true);
      final req =
          await updateArticle(_objKey, title, intro, category, content, imageKey, 'System');
      var result = jsonDecode(req.response);

      print(result);

      window.open("/blog/article/view/${_objKey}", '_blank');
    }
  }

  void onPublishClick(MouseEvent e) async {
    if (isFormValid()) {
      disableSubmit(true);
      final req =
          await publishArticle(_objKey, title, intro, category, content, imageKey, 'System');
      var result = jsonDecode(req.response);

      print(result);
    }
  }

  void onEditorCtrlClick(MouseEvent e) {
    final ctrl = e.matchingTarget;

    if (ctrl is SpanElement) {
      final role = ctrl.dataset['role'];
      
      switch (role) {
        case 'h1':
        case 'h2':
        case 'p':
          document.execCommand('formatBlock', false, role);
          break;
        default:
          document.execCommand(role, false, null);
          break;
      }
    }
  }
}
