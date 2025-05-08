import 'package:flutter/material.dart';
import 'package:auth0_flutter/auth0_flutter.dart';
import 'package:firebase_core/firebase_core.dart';
import 'firebase_options.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import 'constants.dart';

/*

Auth States:
* Full User
* Invitation User
  * View invites
  * Respond to invite
  * Allow to see Private Group/Event
  * Create User if doesn't exist
* No Auth
  * Sign in to invite user or real user
  * View Public Groups
  * View Public Events

What is auth necessary:
* Create Group
* Join Group
* Create Event
* Create Poll
* Create Post
* Post images
* Create profile
* Message

What is not:
* View Groups
* View Events
* View Invitation
* Accept invitation
* Join Event from invite
* Get invitation updates via text/email
* Change your display name?

Invitation:
* Invite from contact (pull phone + name)
* Create "User" from invitation ?
  * What happens if we just store info in the invation itself?
  * Problem: How would you "sign in" to an invitation?
  * Through a link to email/phone?
  * Verify phone number flow?
* They get a text
  * Link to website (sketchy?)
  * How to make less sketchy?
    * Tell them to go to oliveplay.co without clicking anything
  * Sign in flow might look like "type in phone number", then verify
  * Look into Partiful unauth flow
  * How does partiful get away with sending texts without consent?


User:
* auth0_id? nullable
* computed isSimpleUser = auth0_id == null


==========================
|                [login] |
|                        |
|                        |
|                        |
|                        |
|                        |
|                        |
|                        |
|                        |
|                        |
|                        |
==========================

==========================
|              [hi name] |
|                        |
|                        |
|                        |
|                        |
|                        |
|                        |
|                        |
|                        |
|                        |
|                        |
==========================

==========================
|                        |
|                        |
|        OlivePlay       |
|                        |
|         Sign In        |
|          Google        |
|           Apple        |
|            or          |
|                        |
|    View Invitations    |
|      Phone #:          |
==========================

^ Is this a weird UX? 
Pros: People who don't want to create an account won't have to
Cons: Makes the auth screen harder to understand

==========================
|      [invites] [login] |
|                        |
|                        |
|                        |
|                        |
|                        |
|                        |
|                        |
|                        |
|                        |
|                        |
==========================

^ Maybe have them two different buttons like use the `mail` logo
- when not authenticated, then have a green circle to get people to click
- on it more.

==========================
|                        |
|                        |
|                        |
|    View your Invites   |
|        Phone #:        |
|                        |
|                        |
|   Data Charges apply   |
|                        |
|                        |
|                        |
==========================

Just phone number followed by the confirmation should be enough to
create a "auth"

*/

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await Firebase.initializeApp(
    options: DefaultFirebaseOptions.currentPlatform,
  );

  await Constants.initialize();

  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData(
        // This is the theme of your application.
        //
        // TRY THIS: Try running your application with "flutter run". You'll see
        // the application has a purple toolbar. Then, without quitting the app,
        // try changing the seedColor in the colorScheme below to Colors.green
        // and then invoke "hot reload" (save your changes or press the "hot
        // reload" button in a Flutter-supported IDE, or press "r" if you used
        // the command line to start the app).
        //
        // Notice that the counter didn't reset back to zero; the application
        // state is not lost during the reload. To reset the state, use hot
        // restart instead.
        //
        // This works for code too, not just values: Most code changes can be
        // tested with just a hot reload.
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
        useMaterial3: true,
      ),
      home: const MyHomePage(title: 'Flutter Demo Home Page'),
    );
  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({super.key, required this.title});

  // This widget is the home page of your application. It is stateful, meaning
  // that it has a State object (defined below) that contains fields that affect
  // how it looks.

  // This class is the configuration for the state. It holds the values (in this
  // case the title) provided by the parent (in this case the App widget) and
  // used by the build method of the State. Fields in a Widget subclass are
  // always marked "final".

  final String title;

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  int _counter = 0;

  void _incrementCounter() {
    setState(() {
      // This call to setState tells the Flutter framework that something has
      // changed in this State, which causes it to rerun the build method below
      // so that the display can reflect the updated values. If we changed
      // _counter without calling setState(), then the build method would not be
      // called again, and so nothing would appear to happen.
      _counter++;
    });
  }

  late Auth0 auth0;

  @override
  void initState() {
    super.initState();
    auth0 = Auth0('dev-t08v6dtzszs25fnh.us.auth0.com', 'jexEmKJwEYjXLX6EGrqFjLlUZhwYcxEk');
  }

  Credentials? _credentials;

  @override
  Widget build(BuildContext context) {
    // This method is rerun every time setState is called, for instance as done
    // by the _incrementCounter method above.

    if (_credentials == null) {
      return Scaffold(
        body: Center(
          child: ElevatedButton(
            onPressed: () async {
              // Use a Universal Link callback URL on iOS 17.4+ / macOS 14.4+
              // useHTTPS is ignored on Android
              final credentials =
                  await auth0.webAuthentication().login(useHTTPS: true);

              setState(() {
                _credentials = credentials;
              });
            },
            child: const Text("Log in")
          ),
        ),
      );
    }

    // The Flutter framework has been optimized to make rerunning build methods
    // fast, so that you can just rebuild anything that needs updating rather
    // than having to individually change instances of widgets.
    return Scaffold(
      appBar: AppBar(
        // TRY THIS: Try changing the color here to a specific color (to
        // Colors.amber, perhaps?) and trigger a hot reload to see the AppBar
        // change color while the other colors stay the same.
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        // Here we take the value from the MyHomePage object that was created by
        // the App.build method, and use it to set our appbar title.
        title: Text(widget.title),
      ),
      body: Center(
        // Center is a layout widget. It takes a single child and positions it
        // in the middle of the parent.
        child: Column(
          // Column is also a layout widget. It takes a list of children and
          // arranges them vertically. By default, it sizes itself to fit its
          // children horizontally, and tries to be as tall as its parent.
          //
          // Column has various properties to control how it sizes itself and
          // how it positions its children. Here we use mainAxisAlignment to
          // center the children vertically; the main axis here is the vertical
          // axis because Columns are vertical (the cross axis would be
          // horizontal).
          //
          // TRY THIS: Invoke "debug painting" (choose the "Toggle Debug Paint"
          // action in the IDE, or press "p" in the console), to see the
          // wireframe for each widget.
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            const Text(
              'You have pushed the button this many times:',
            ),
            Text(
              '$_counter',
              style: Theme.of(context).textTheme.headlineMedium,
            ),
          ],
        ),
      ),
      floatingActionButtonLocation: FloatingActionButtonLocation.centerFloat,
      persistentFooterButtons: [
        ElevatedButton(
          onPressed: () async {
            final url = Uri.parse('${Constants.apiUrl}/ping');
            try {
              final response = await http.post(
                url,
                headers: {'Content-Type': 'application/json'},
                body: jsonEncode({'message': 'hi'}),
              );
              ScaffoldMessenger.of(context).showSnackBar(
                SnackBar(
                  content: Text('Response: ${response.body}'),
                  duration: const Duration(seconds: 3),
                ),
              );
            } catch (e) {
              ScaffoldMessenger.of(context).showSnackBar(
                SnackBar(
                  content: Text('Error: $e'),
                  backgroundColor: Colors.red,
                  duration: const Duration(seconds: 3),
                ),
              );
            }
          },
          child: const Text('Ping API'),
        ),
        ElevatedButton(
          onPressed: () async {
            // Use a Universal Link logout URL on iOS 17.4+ / macOS 14.4+
            // useHTTPS is ignored on Android
            await auth0.webAuthentication().logout(useHTTPS: true);
  
            setState(() {
              _credentials = null;
            });
          },
          child: const Text('Logout'),
        ),
      ],
      bottomNavigationBar: BottomAppBar(
        child: Padding(
          padding: const EdgeInsets.all(8.0),
          child: Column(
            mainAxisSize: MainAxisSize.min,
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              if (_credentials?.user.name != null)
                Text('Name: ${_credentials?.user.name}'),
              if (_credentials?.user.email != null)  
                Text('Email: ${_credentials?.user.email}'),
              if (_credentials?.user.nickname != null)
                Text('Nickname: ${_credentials?.user.nickname}'),
            ],
          ),
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: _incrementCounter,
        tooltip: 'Increment',
        child: const Icon(Icons.add),
      ), // This trailing comma makes auto-formatting nicer for build methods.
    );
  }
}
