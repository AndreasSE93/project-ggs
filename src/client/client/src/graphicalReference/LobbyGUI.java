package graphicalReference;

import java.awt.*;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.text.DateFormat.Field;
import java.util.ArrayList;

import javax.swing.*;

import org.json.JSONException;

import packageManaging.LobbyClientMessage;
import packageManaging.LobbyHandler;
import packageManaging.Message;
import packageManaging.MessageHandler;

import clientNetworking.NetManager;

public class LobbyGUI extends GUI {
		JPanel chatt;
		public JTextField field;
		public JTextPane chatPanel = new JTextPane();
		
		public NetManager NM;

	public void render( ArrayList<String> players,  ArrayList<String> games, NetManager m) {
		
		this.NM = m;
		
		JFrame lobby = new JFrame();
		lobby.setLayout(new BorderLayout());
		lobby.getContentPane().setBackground(Color.LIGHT_GRAY);
		lobby.setSize(800,600);
        lobby.setLocationRelativeTo(null);
        lobby.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        
        
        JPanel topPanel = new JPanel();
        JPanel bottomPanel = new JPanel();
        JPanel eastTopPanel = new JPanel();
        JPanel westTopPanel = new JPanel();
        topPanel.setLayout(new BorderLayout());
        eastTopPanel.setLayout(new BorderLayout());
        westTopPanel.setLayout(new BorderLayout());
        eastTopPanel.setBackground(Color.BLACK);
        westTopPanel.setBackground(Color.RED);
        makeList(eastTopPanel, players,"gogogog");
        makeList(westTopPanel, games,"gogogog");
        topPanel.add(westTopPanel, BorderLayout.LINE_END);
        topPanel.add(eastTopPanel, BorderLayout.LINE_START);
        topPanel.setBackground(Color.GREEN);
        Dimension preferredSize = new Dimension(400,300);
        topPanel.setPreferredSize(preferredSize);
        lobby.add(topPanel, BorderLayout.NORTH);
        lobby.add(bottomPanel, BorderLayout.SOUTH);
        JPanel chatt = new JPanel();
        chatStart(chatt);
        lobby.add(chatt);
        
        lobby.setVisible(true);
		
	}

	@Override
	public void actionPerformed(ActionEvent e) {
		String str = this.field.getText();
		
		MessageHandler l = (MessageHandler)this.NM.handler;
		Message mess = new Message(str, "Petter", 10);
		String sender = "";
		try {
			sender = l.encode(mess);
		} catch (JSONException e1) {
			e1.printStackTrace();
		}
		
		this.NM.send(sender);
		
	}
	
	public static void makeList(JPanel panel ,ArrayList<String> L, String nameButton){
		JButton game = new JButton(nameButton);
		final JList<String> jList = getJList(L);
		Dimension preferredSize = new Dimension(200 , 200);
		jList.setPreferredSize(preferredSize);
		panel.add(game, BorderLayout.SOUTH);
		panel.add(jList, BorderLayout.NORTH);
		game.addActionListener(new ActionListener(){
        	public void actionPerformed(ActionEvent ae){
        		String selectedItem = (String) jList.getSelectedValue();
        		System.out.println("Clicked" + selectedItem);
        	}
		});
	}
	
	public void chatStart(JPanel j){
		
		//String userName = (String)JOptionPane.showInputDialog("Write username!");
        j.setLayout(new BorderLayout());
        this.field = new JTextField(20);
        field.addActionListener(this);
        
        chatPanel.setVisible(true);
        chatPanel.setEditable(false);
        chatPanel.setForeground(Color.BLUE);
        j.add(chatPanel, "North");
        j.add(field, "South");
        
        
		
        
		

	}
	
	

	@Override
	public void render() {
		// TODO Auto-generated method stub
		
	}
	public static JList<String> getJList(ArrayList<String> list){
		int size = list.size();
		final DefaultListModel<String> model = new DefaultListModel<String>();
		for (int i=0; i<size; i++){
			model.addElement(list.get(i));
		}
		return new JList<String>(model);
	}
}
