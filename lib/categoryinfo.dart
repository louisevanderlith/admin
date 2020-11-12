import 'dart:html';

import 'package:mango_ui/keys.dart';

class CategoryInfoForm {
  TextInputElement txtName;
  TextInputElement txtText;
  TextInputElement txtDescription;
  TextInputElement txtClient;
  FileUploadInputElement uplImage;

  CategoryInfoForm() {
    txtText = querySelector("#txtInfoText");
    txtName = querySelector("#txtInfoName");
    txtDescription = querySelector("#txtInfoDescription");
    txtClient = querySelector("#txtInfoClient");
    uplImage = querySelector("#uplInfoImageImg");
  }

  String get name {
    return txtName.value;
  }

  String get text {
    return txtText.value;
  }

  String get description {
    return txtDescription.value;
  }

  String get client {
    return txtClient.value;
  }

  Key get image {
    return new Key(uplImage.dataset["id"]);
  }
}
