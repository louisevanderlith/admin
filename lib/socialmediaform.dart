import 'dart:html';

import 'package:mango_ui/bodies/sociallink.dart';
import 'package:mango_ui/formstate.dart';
import 'package:mango_ui/trustvalidator.dart';
import 'models/socialmediaitem.dart';

class SocialmediaForm extends FormState {
  SocialmediaForm(String idElem, String btnSubmit, String btnAdd)
      : super(idElem, btnSubmit) {
        findSocials();
    querySelector(btnAdd).onClick.listen(onAddClick);
  }

  void onAddClick(MouseEvent e) {
    addSocial();
  }

  List<Sociallink> get items {
    return findSocials();
  }

  List<Sociallink> findSocials() {
    var hasSocial = false;
    var result = new List<Sociallink>();
    var indx = 0;

    do {
      var social =
          new SocialmediaItem('#txtSocialIcon${indx}', '#txtSocialURL${indx}');
      hasSocial = social.loaded();

      if (hasSocial) {
        result.add(social.toDTO());
      }

      indx++;
    } while (hasSocial);

    return result;
  }

  void addSocial() {
    //add elements
    var indx = items.length;

    var schema = buildElement(indx);
    form.children.add(schema);

 /*   var item =
        new SocialmediaItem('#txtSocialIcon${indx}', '#txtSocialURL${indx}');

    _items.add(item.toDTO());*/
  }

  //returns HTML for a SocialMedia Item
  Element buildElement(int index) {
    var schema = '''
    <div class="card">
                <header class="card-header">
                    <a href="#" data-liID="liSocial${index}" class="card-header-icon" aria-label="more options">
                        <span class="icon">
                            <i class="fas fa-close" aria-hidden="true"></i>
                        </span>
                    </a>
                </header>
                <div class="card-content">
                    <div class="content">
                        <div class="field">
                            <div class="control">
                                <label class="label" for="txtSocialURL${index}">URL
                                    <input class="input" type="text" min-length="3" id="txtSocialURL${index}" required
                                        value="" />
                                    <p class="help is-danger"></p>
                            </div>
                        </div>
                        <div class="field">
                            <div class="control">
                                <label class="label" for="txtSocialIcon${index}">Icon
                                    <input class="input" type="text" id="txtSocialIcon${index}" required
                                        value="" />
                                    <p class="help is-danger"></p>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        ''';

    return new Element.html(schema, validator: new TrustedNodeValidator());
  }
}
