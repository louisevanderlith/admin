import 'dart:html';

import 'package:Admin.APP/formstate.dart';
import 'package:Admin.APP/services/blogapi.dart';

class BlogForm extends FormState {
  TextInputElement _title;
  TextAreaElement _content;
  FileUploadInputElement _headImage;

  BlogForm(
      String idElem, String titleElem, String contentElem, String submitBtn)
      : super(idElem, submitBtn) {
    _title = querySelector(titleElem);
    _content = querySelector(contentElem);

    querySelector(submitBtn).onClick.listen(onSubmitClick);
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

  void onSubmitClick(MouseEvent e) {
    if (isFormValid()) {
      disableSubmit(true);
      updateArticle(title, content, imageKey, 'System');
    }
  }
}
