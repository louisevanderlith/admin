import 'dart:html';

import 'package:Admin.APP/models/mapitem.dart';
import 'package:mango_artifact/uploadapi.dart';
import 'package:mango_ui/keys.dart';

class BasicSiteForm {
  FormElement form;
  TextInputElement txtTitle;
  TextAreaElement txtDescription;
  FileUploadInputElement uplProfileImg;

  List<MapItem> lstEndpoints;
  List<MapItem> lstCodes;
  List<MapItem> lstTerms;

  BasicSiteForm() {
    form = querySelector("frmBasicSite");
    txtTitle = querySelector("txtTitle");
    txtDescription = querySelector("txtDescription");
    uplProfileImg = querySelector("uplProfileImg");

    uplProfileImg.onChange.listen(uploadFile);
  }

  String get title {
    return txtTitle.text;
  }

  String get description {
    return txtDescription.text;
  }

  Key get imageKey {
    return new Key(uplProfileImg.dataset['id']);
  }

  Map<String, String> get endpoints {
    var result = new Map<String, String>();

    result.addEntries(lstEndpoints.map((e) => new MapEntry(e.name, e.value)));

    return result;
  }

  Map<String, String> get codes {
    var result = new Map<String, String>();
    result.addEntries(lstCodes.map((e) => new MapEntry(e.name, e.value)));

    return result;
  }

  Map<String, String> get terms {
    var result = new Map<String, String>();
    result.addEntries(lstTerms.map((e) => new MapEntry(e.name, e.value)));

    return result;
  }
}
