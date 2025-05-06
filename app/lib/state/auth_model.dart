import 'package:flutter/foundation.dart';
import 'package:auth0_flutter/auth0_flutter.dart';

class AuthModel extends ChangeNotifier {
  /// Internal, private state of the credentials.
  Credentials? _credentials;

  /// Check if user is authenticated
  bool get isAuthenticated => _credentials != null;

  /// Get the user's name, or empty string if not available
  String get userName => _credentials?.user.name ?? '';
  
  /// Get the user's email, or empty string if not available
  String get userEmail => _credentials?.user.email ?? '';
  
  /// Get the user's nickname, or empty string if not available
  String get userNickname => _credentials?.user.nickname ?? '';

  void setCredentials(Credentials credentials) {
    _credentials = credentials;
    notifyListeners();
  }

  void clearCredentials() {
    _credentials = null;
    notifyListeners();
  }

  /// Adds [item] to cart. This and [removeAll] are the only ways to modify the
  // void add(Item item) {
  //   _items.add(item);
  //   // This call tells the widgets that are listening to this model to rebuild.
  //   notifyListeners();
  // }

  // /// Removes all items from the cart.
  // void removeAll() {
  //   _items.clear();
  //   // This call tells the widgets that are listening to this model to rebuild.
  //   notifyListeners();
  // }
}