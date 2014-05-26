package packageManaging;

import org.json.JSONArray;

import org.json.JSONException;
import org.json.JSONObject;

import snake.SnakePlayer;
import snake.SnakeServerMessage;
import snake.SnakeUserMessage;

/* 							 CLASS SPECIFICATION 
 *        Encoders (Client -> Server) and decoders (Server -> client) for JSON messages sent by and to the server
 *        Each message will have a specified ID to know what type of message is being sent and recieved to server.
 *        
 *        ID Specification list:
 *        
 *        Class: InitializeClientMessage:                    PacketID 99
 *        Class: Message:                                    PacketID 100
 *        Class: CreateGameMessage:                          PacketID 101
 *        Class: JoinMessage                                 PacketID 102
 *        Class: LobbyServerMessage, RefreshMessage          PacketID 103
 *        Class: TiarUserMessage                             PacketID 201
 *        Class: StartableMessage							 PacketID 202
 *        Class: StartedMessage								 PacketID 203
 *        Class: KickMessage                                 PacketID 404
 */

public class Encoder {


	public String encode(Message mess) throws JSONException {

		JSONObject obj = new JSONObject();

		obj.put("message", mess.message);
		obj.put("PacketID", mess.id);
		obj.put("user", mess.user);

		String message = obj.toString();
		return message;

	}

	public Message chattMessageDecode(String enc) throws JSONException {

		JSONObject obj = new JSONObject(enc);

		Message chatMessage = new Message();
		chatMessage.message = obj.getString("message");
		chatMessage.user = obj.getString("user");

		return chatMessage;

	}

	public String encode(CreateGameMessage mess) throws JSONException {

		JSONObject obj = new JSONObject();

		obj.put("GameName", mess.gameName);
		obj.put("PacketID", mess.PacketID);
		obj.put("MaxSize", mess.maxSize);
		obj.put("RoomName", mess.roomName);

		String message = obj.toString();
		return message;

	}

	public CreateGameMessage createGameMessageDecode(String mess)
			throws JSONException {
		JSONObject obj = new JSONObject(mess);

		CreateGameMessage lsm = new CreateGameMessage();

		lsm.gameName = obj.getString("gameName");

		return lsm;
	}

	public HostRoom hostRoomDecode(String mess) throws JSONException {
		JSONObject obj = new JSONObject(mess);

		HostRoom hr = new HostRoom();

		hr.Player = obj.getInt("Player");

		return hr;
	}

	public String encode(InitializeClientMessage mess) throws JSONException {

		JSONObject obj = new JSONObject();

		obj.put("PacketID", mess.PacketID);
		obj.put("UserName", mess.userName);

		String message = obj.toString();
		return message;

	}

	public String encode(JoinMessage mess) throws JSONException {

		JSONObject obj = new JSONObject();

		obj.put("PacketID", mess.PacketID);
		obj.put("RoomID", mess.RoomID);

		String message = obj.toString();
		return message;

	}

	public JoinMessage joinMessageDecode(String mess) throws JSONException {
		JSONObject obj = new JSONObject(mess);
		System.out.println(mess);

		JoinMessage lsm = new JoinMessage();
		JSONObject i = obj.getJSONObject("hostRoom");

		lsm.RoomID = i.getInt("RoomID");
		lsm.GameType = i.getString("GameType");
		lsm.GameName = i.getString("GameName");

		return lsm;
	}

	public KickMessage kickMessageDecode(String mess) throws JSONException {

		KickMessage km = new KickMessage();

		return km;

	}

	public String encode(KickMessage mess) throws JSONException {

		JSONObject obj = new JSONObject();

		obj.put("PacketID", mess.PacketID);

		String message = obj.toString();
		return message;

	}

	public LobbyServerMessage lobbyServerMessageDecode(String mess)
			throws JSONException {
		JSONObject obj = new JSONObject(mess);

		LobbyServerMessage lsm = new LobbyServerMessage();

		JSONArray JArr = obj.getJSONArray("UserList");
		for (int i = 0; i < JArr.length(); i++) {
			JSONObject l = JArr.getJSONObject(i);
			HostRoom h = new HostRoom(l.getInt("RoomID"), l.getInt("MaxSize"),
					l.getInt("ClientCount"), l.getString("RoomName"),
					l.getString("GameName"));
			lsm.UserList.add(h);
		}
		JSONArray JArr2 = obj.getJSONArray("GameHost");
		for (int i = 0; i < JArr2.length(); i++) {
			lsm.GameHost.add(JArr2.getString(i));
		}

		return lsm;
	}

	public String encode(RefreshMessage mess) throws JSONException {
		JSONObject obj = new JSONObject();

		obj.put("PacketID", mess.PacketID);

		String message = obj.toString();
		return message;

	}

	public StartableMessage startebleMessageDecode(String mess)
			throws JSONException {
		JSONObject obj = new JSONObject(mess);

		StartableMessage tum = new StartableMessage();

		tum.isStartable = obj.getBoolean("IsStartable");

		return tum;
	}

	public String encode(StartedMessage mess) throws JSONException {
		JSONObject obj = new JSONObject();

		obj.put("PacketID", mess.PacketID);

		String message = obj.toString();
		return message;

	}

	public StartedMessage startedMessageDecode(String mess)
			throws JSONException {
		JSONObject obj = new JSONObject(mess);

		StartedMessage tum = new StartedMessage();

		tum.started = obj.getBoolean("Started");
		tum.playerID = obj.getInt("Player");

		return tum;
	}



	public String encode(TiarUserMessage mess) throws JSONException {

		JSONObject obj = new JSONObject();

		obj.put("PacketID", mess.PacketID);
		obj.put("Move", mess.Move);
		obj.put("Player", mess.player);

		String message = obj.toString();
		return message;

	}

	public TiarUserMessage tiarUserMessageDecode(String enc)
			throws JSONException {

		JSONObject obj = new JSONObject(enc);

		TiarUserMessage TiarMessage = new TiarUserMessage();
		JSONArray JArray = obj.getJSONArray(("GameBoard"));
		int[] array = new int[9];
		for (int i = 0; i < array.length; i++) {
			array[i] = JArray.getInt(i);
		}
		TiarMessage.Gameboard = array;
		TiarMessage.HasWon = obj.getInt("HasWon");
		TiarMessage.IsDraw = obj.getInt("IsDraw");
		TiarMessage.isValid = obj.getInt("IsValid");

		return TiarMessage;

	}

	public String encode(SnakeUserMessage mess) throws JSONException {

		JSONObject obj = new JSONObject();

		obj.put("PacketID", mess.PacketID);
		obj.put("PlayerID", mess.PlayerID);
		obj.put("KeyEvent", mess.KeyEvent);

		String message = obj.toString();
		return message;

	}

	public SnakeServerMessage snakeServerMessageDecode(String mess)
			throws JSONException {
		JSONObject obj = new JSONObject(mess);

		double[] p;
		JSONArray JArr = obj.getJSONArray("PlayerArray");
		SnakeServerMessage SSM = new SnakeServerMessage(JArr.length());
		for (int i = 0; i < JArr.length(); i++) {
			p = new double[10];
			JSONObject l = JArr.getJSONObject(i);
			JSONArray JArr2 = l.getJSONArray("Coordinates");

			for (int d = 0; d < p.length; d++) {
				p[d] = JArr2.getDouble(d);

			}
			SnakePlayer sp = new SnakePlayer(l.getInt("PlayerID"),
					l.getDouble("PosX"), l.getDouble("PosY"),
					l.getBoolean("Alive"), l.getInt("Score"), p);
			sp.PlayerName = l.getString("PlayerName");
			// System.out.println(sp.PlayerName);
			SSM.hasWon = obj.getBoolean("HasWon");
			SSM.winnerName = obj.getString("WinnerName");
			SSM.Players[i] = sp;
			SSM.clearBoard = obj.getBoolean("ClearBoard");
		}

		return SSM;
	}
}
