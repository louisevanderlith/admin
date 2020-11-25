import 'dart:html';

import 'package:mango_artifact/uploadapi.dart';
import 'package:mango_ui/keys.dart';

class CategoryInfoForm {
  TextInputElement txtName;
  TextInputElement txtText;
  TextInputElement txtDescription;
  UrlInputElement txtPageURL;
  SelectElement cboBaseCategory;
  TextInputElement txtClient;
  FileUploadInputElement uplImage;

  CategoryInfoForm() {
    txtText = querySelector("#txtInfoText");
    txtName = querySelector("#txtInfoName");
    txtDescription = querySelector("#txtInfoDescription");
    txtPageURL = querySelector("#txtPageURL");
    cboBaseCategory = querySelector("#cboBaseCategory");
    txtClient = querySelector("#txtInfoClient");
    uplImage = querySelector("#uplInfoImageImg");

    uplImage.onChange.listen(uploadFile);
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

  String get pageurl {
    return txtPageURL.value;
  }

  num get basecategory {
    return num.parse(cboBaseCategory.value);
  }

  String get client {
    return txtClient.value;
  }

  Key get image {
    return new Key(uplImage.dataset["id"]);
  }
}
