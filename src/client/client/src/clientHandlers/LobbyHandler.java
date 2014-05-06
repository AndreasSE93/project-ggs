package clientHandlers;

import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.io.IOException;

import org.json.JSONException;
import org.json.JSONObject;

import packageManaging.LobbyMessageEncoder;
import packageManaging.LobbyServerMessage;
import packageManaging.Message;
import packageManaging.ChatMessageEncoder;
import graphicalReference.ChatGUI;
import graphicalReference.LobbyGUI;

import clientNetworking.NetManager;

public class LobbyHandler implements HandlerInterface,
		ActionListener {
	LobbyGUI lg;
	ChatGUI cg;
	LobbyMessageEncoder lme = new LobbyMessageEncoder();
	ChatMessageEncoder cme = new ChatMessageEncoder();
	NetManager network;
	
	public LobbyHandler(NetManager net){
		this.network = net;
	}
	
	public void initHandler() {

		String firstcall;
		LobbyServerMessage LM = null;
		try {
			firstcall = network.receiveMessage();

			System.out.println(firstcall + "\n");

			LM = lme.decode(firstcall);

			System.out.println(LM + "\n");
		} catch (IOException | JSONException e) {
			e.printStackTrace();
		}

		if (LM != null) {
			lg = new LobbyGUI();
			lg.render(LM.UserList, LM.GameHost);
			lg.chatgui.field.addActionListener(this);

		} else {
			System.out.println("Connection failed");
		}

		while (true) {
			try {
				String mess = recieveMessage();
			
				int id = retrieveId(mess);
				System.out.println(id);

				decodeAndRender(id, mess);
				
			} catch (IOException | JSONException e) {

				e.printStackTrace();
			}

		}

	}

	public void actionPerformed(ActionEvent e) {
		String command = e.getActionCommand();
		String JSONtext = "";
		switch (command) {
		case "chatmessage":  //JTextField
			try {
				JSONtext = cme.encode(new Message(lg.chatgui.field.getText(), lg.chatgui.userName));
				lg.chatgui.field.setText("");
			} catch (JSONException e1) {
				e1.printStackTrace();
			}

			break;

		default:
			JSONtext = "";

		}
		sendMessage(JSONtext);

	}

	public void decodeAndRender(int id, String message) throws JSONException {
		switch (id) {
		case 100: // Message.java
			ChatMessageEncoder mh = new ChatMessageEncoder();
			Message chatMessage = mh.decode(message);
			lg.chatgui.chatUpdate(chatMessage.message, chatMessage.user);
			break;

		case 101: // Create session
			break;

		case 102: // Join session
			break;

		case 103: // LobbyClientMessage.java/LobbyServerMessage.java
			break;
		default:
			System.out.println("something went wrong!");
		}

	}

	public void sendMessage(String message) {
		network.send(message);
	}

	public String recieveMessage() throws IOException {
		return network.receiveMessage();

	}

	public int retrieveId(String mess) {

		try {
			JSONObject obj = new JSONObject(mess);
			return obj.getInt("PacketID");
		} catch (JSONException e) {
			
		}
		return 0;

	}
}
