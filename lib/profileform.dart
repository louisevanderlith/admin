import 'dart:convert';
import 'dart:html';
import 'package:mango_ui/bodies/header.dart';
import 'package:mango_ui/bodies/key.dart';
import 'package:mango_ui/bodies/portfolio.dart';
import 'package:mango_ui/bodies/profile.dart';
import 'package:mango_ui/bodies/sociallink.dart';
import 'package:mango_ui/formstate.dart';

import 'package:mango_ui/services/profileapi.dart';
import 'headerform.dart';
import 'models/headeritem.dart';
import 'portfolioform.dart';
import 'models/portfolioitem.dart';
import 'socialmediaform.dart';
import 'models/socialmediaitem.dart';
import 'package:mango_ui/services/uploadapi.dart';

class ProfileForm extends FormState {
  Key _objKey;
  TextInputElement _name;
  TextAreaElement _description;
  EmailInputElement _email;
  TelephoneInputElement _phone;
  TextInputElement _url;
  TextInputElement _gtag;
  FileUploadInputElement _image;

  HeaderForm _headers;
  PortfolioForm _portfolios;
  SocialmediaForm _socialmedia;

  ParagraphElement _error;

  ProfileForm(
      String idElem,
      Key objKey,
      String nameElem,
      String descElem,
      String emailElem,
      String phoneElem,
      String urlElem,
      String gtagElem,
      String imageElem,
      String frmHeader,
      String addHeader,
      String frmPortfolio,
      String addPortfolio,
      String frmSocialmedia,
      String addSocial,
      String submitBtn)
      : super(idElem, submitBtn) {
    _objKey = objKey;
    _name = querySelector(nameElem);
    _description = querySelector(descElem);
    _email = querySelector(emailElem);
    _phone = querySelector(phoneElem);
    _url = querySelector(urlElem);
    _image = querySelector(imageElem);
    _gtag = querySelector(gtagElem);
    _error = querySelector("${idElem}Err");

    _headers = new HeaderForm(frmHeader, submitBtn, addHeader);
    _socialmedia = new SocialmediaForm(frmSocialmedia, submitBtn, addSocial);
    _portfolios = new PortfolioForm(frmPortfolio, submitBtn, addPortfolio);

    querySelector(submitBtn).onClick.listen(onSend);

    _image.onChange.listen(uploadFile);
  }

  String get name {
    return _name.value;
  }

  String get description {
    return _description.value;
  }

  String get email {
    return _email.value;
  }

  String get phone {
    return _phone.value;
  }

  String get url {
    return _url.value;
  }

  String get gtag {
    return _gtag.value;
  }
  
  Key get imageKey {
    return new Key(_image.dataset["id"]);
  }

  List<Portfolio> get portfolioItems {
    return _portfolios.items;
  }

  List<Sociallink> get socialmediaItems {
    return _socialmedia.items;
  }

  List<Header> get headerItems {
    return _headers.items;
  }

  void onSend(MouseEvent e) async {
    if (isFormValid()) {
      disableSubmit(true);

      final data = new Profile(name, description, email, phone, url, gtag,
          imageKey, portfolioItems, socialmediaItems, headerItems);
      var req = await updateProfile(_objKey, data);

      final resp = jsonDecode(req.response);

      if (req.status == 200) {
        window.alert(resp['Data']);
      } else {
        _error.text = resp['Error'];
      }
    }
  }
}
