import '../formstate.dart';
import 'socialmediaitem.dart';

class SocialmediaForm extends FormState {
  List<SocialmediaItem> _items;

  SocialmediaForm(String idElem, String btnSubmit) : super(idElem, btnSubmit) {
    _items = findSocials(btnSubmit);
  }

  List<SocialmediaItem> get items {
    return _items;
  }

  List<SocialmediaItem> findSocials(String btnSubmit) {
    var hasSocial = false;
    var result = new List<SocialmediaItem>();
    var indx = 0;

    do {
      var social =
          new SocialmediaItem('#txtSocialIcon${indx}', '#txtSocialURL${indx}');
      hasSocial = social.loaded();

      if (hasSocial) {
        result.add(social);
      }

      indx++;
    } while (hasSocial);

    return result;
  }
}
