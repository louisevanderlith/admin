import 'dart:convert';
import 'dart:html';

import 'package:Admin.APP/formstate.dart';
import 'package:Admin.APP/services/blogapi.dart';
import 'package:Admin.APP/services/uploadapi.dart';

class BlogForm extends FormState {
  String _objKey;
  TextInputElement _title;
  TextAreaElement _content;
  FileUploadInputElement _headImage;

  BlogForm(String idElem, String objKey, String titleElem, String contentElem,
      String imageElem, String previewBtn, String publishBtn, String submitBtn)
      : super(idElem, submitBtn) {
    _objKey = objKey;
    _title = querySelector(titleElem);
    _content = querySelector(contentElem);
    _headImage = querySelector(imageElem);

    querySelector(submitBtn).onClick.listen(onSubmitClick);
    querySelector(previewBtn).onClick.listen(onPreviewClick);
    querySelector(publishBtn).onClick.listen(onPublishClick);

    _headImage.onChange.listen(uploadFile);
  }

  String get title {
    return _title.value;
  }

  String get content {
    return _content.value;
  }

  String get imageKey {
    return _headImage.dataset['id'];
  }

  void onSubmitClick(MouseEvent e) async {
    if (isFormValid()) {
      disableSubmit(true);
      final req =
          await updateArticle(_objKey, title, content, imageKey, 'System');
      var result = jsonDecode(req.response);

      print(result);
    }
  }

  void onPreviewClick(MouseEvent e) async {
    if (isFormValid()) {
      disableSubmit(true);
      final req =
          await updateArticle(_objKey, title, content, imageKey, 'System');
      var result = jsonDecode(req.response);

      print(result);

      window.open("/blog/view/${_objKey}", '_blank');
    }
  }

  void onPublishClick(MouseEvent e) async {
    if (isFormValid()) {
      disableSubmit(true);
      final req =
          await publishArticle(_objKey, title, content, imageKey, 'System');
      var result = jsonDecode(req.response);

      print(result);
    }
  }
}
